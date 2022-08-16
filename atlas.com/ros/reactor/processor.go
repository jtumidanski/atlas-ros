package reactor

import (
	"atlas-ros/model"
	"atlas-ros/reactor/script"
	registry2 "atlas-ros/reactor/script/registry"
	"atlas-ros/reactor/statistics"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func DeferCreate(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, classification uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) {
	return emitCreateReactor(l, span)
}

func Create(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, classification uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) (*Model, error) {
	return func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) (*Model, error) {
		s, err := statistics.GetCache().GetFile(reactorId)
		if err != nil {
			l.WithError(err).Errorf("Unable to get reactor statistics from WZ")
			return nil, err
		}
		m := GetRegistry().Create(worldId, channelId, mapId, reactorId, name, state, x, y, delay, direction, *s)
		emitCreated(l, span)(worldId, channelId, mapId, m.Id())
		return &m, nil
	}
}

func GetInMap(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) []Model {
	return func(worldId byte, channelId byte, mapId uint32) []Model {
		return GetRegistry().GetInMap(worldId, channelId, mapId)
	}
}

func ByIdProvider(id uint32) model.Provider[Model] {
	return func() (Model, error) {
		return GetRegistry().Get(id)
	}
}

func InMapProvider(worldId byte, channelId byte, mapId uint32) model.SliceProvider[Model] {
	return func() ([]Model, error) {
		return GetRegistry().GetInMap(worldId, channelId, mapId), nil
	}
}

func ByNameFilter(name string) model.PreciselyOneFilter[Model] {
	return func(rs []Model) (Model, error) {
		for _, r := range rs {
			if r.Name() == name {
				return r, nil
			}
		}
		return Model{}, errors.New("no match")
	}
}

func ByNameInMapProvider(worldId byte, channelId byte, mapId uint32, name string) model.Provider[Model] {
	return model.SliceProviderToProviderAdapter[Model](InMapProvider(worldId, channelId, mapId), ByNameFilter(name))
}

func ByClassificationFilter(classification uint32) model.PreciselyOneFilter[Model] {
	return func(rs []Model) (Model, error) {
		for _, r := range rs {
			if r.Classification() == classification {
				return r, nil
			}
		}
		return Model{}, errors.New("no match")
	}
}

func ByClassificationInMapProvider(worldId byte, channelId byte, mapId uint32, classification uint32) model.Provider[Model] {
	return model.SliceProviderToProviderAdapter[Model](InMapProvider(worldId, channelId, mapId), ByClassificationFilter(classification))
}

func GetById(id uint32) (Model, error) {
	return ByIdProvider(id)()
}

func GetByNameInMap(worldId byte, channelId byte, mapId uint32, name string) (Model, error) {
	return ByNameInMapProvider(worldId, channelId, mapId, name)()
}

func GetByClassificationInMap(worldId byte, channelId byte, mapId uint32, classification uint32) (Model, error) {
	return ByClassificationInMapProvider(worldId, channelId, mapId, classification)()
}

func Hit(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB) func(id uint32, characterId uint32, stance uint16, skillId uint32) error {
	return func(id uint32, characterId uint32, stance uint16, skillId uint32) error {
		r, err := GetById(id)
		if err != nil {
			return err
		}
		if !r.Active() {
			return nil
		}
		clearTimeout(l)(r)

		err = performScriptAction(l, span, db)(characterId, script.InvokeHit)(r)
		if err != nil {
			return err
		}

		if r.Type() < 999 && r.Type() != -1 {
			if !(r.Type() == 2 && (stance == 0 || stance == 2)) {
				for i := byte(0); i < r.StateSize(); i++ {
					activeSkills := r.ActiveSkills(i)
					if len(activeSkills) > 0 {
						if !contains(activeSkills, skillId) {
							continue
						}
					}
					r, err = GetRegistry().Update(id, advanceState(i))
					if err != nil {
						return err
					}
					if r.NextState(i) == -1 {
						if r.Type() < 100 {
							if r.Delay() > 0 {
								destroy(l, span)(r)
							} else {
								trigger(l, span)(r, stance)
							}
						} else {
							trigger(l, span)(r, stance)
						}
						err = performScriptAction(l, span, db)(characterId, script.InvokeAct)(r)
						if err != nil {
							return err
						}
					} else {
						trigger(l, span)(r, stance)
						if r.State() == r.NextState(i) {
							err = performScriptAction(l, span, db)(characterId, script.InvokeAct)(r)
							if err != nil {
								return err
							}
						}
						r, err = GetRegistry().Update(id, shouldCollect(true))
						if err != nil {
							return err
						}
						refreshTimeout(l, span)(r)
						if r.Type() == 100 {
							searchItem(l)(r)
						}
					}
				}
			}
		} else {
			r, err = GetRegistry().Update(id, incrementState(), shouldCollect(true))
			if err != nil {
				return err
			}
			trigger(l, span)(r, stance)
			if r.Classification() != 9980000 && r.Classification() != 9980001 {
				err = performScriptAction(l, span, db)(characterId, script.InvokeAct)(r)
				if err != nil {
					return err
				}
			}
			refreshTimeout(l, span)(r)
			if r.Type() == 100 {
				searchItem(l)(r)
			}
		}
		return nil
	}
}

func trigger(l logrus.FieldLogger, span opentracing.Span) func(r Model, stance uint16) {
	return func(r Model, stance uint16) {
		emitTriggered(l, span)(r.WorldId(), r.ChannelId(), r.MapId(), r.Id(), stance)
	}
}

func destroy(l logrus.FieldLogger, span opentracing.Span) func(r Model) {
	return func(r Model) {
		dr, err := GetRegistry().Destroy(r.Id())
		if err != nil {
			l.WithError(err).Errorf("Unable to destroy reactor %d.", r.Id())
			return
		}
		clearTimeout(l)(dr)
		emitDestroyed(l, span)(dr.WorldId(), dr.ChannelId(), dr.MapId(), dr.Id())
		go respawn(l, span)(dr)
	}
}

func respawn(l logrus.FieldLogger, span opentracing.Span) func(r Model) {
	return func(r Model) {
		time.Sleep(time.Duration(r.Delay()) * time.Millisecond)

		GetRegistry().Remove(r.Id())

		_, err := Create(l, span)(r.WorldId(), r.ChannelId(), r.MapId(), r.Classification(), r.Name(), 0, r.X(), r.Y(), r.Delay(), r.FacingDirection())
		if err != nil {
			l.WithError(err).Errorf("Unable to respawn reactor %d at location %d,%d in map %d.", r.Classification(), r.X(), r.Y(), r.MapId())
		}
	}
}

func clearTimeout(_ logrus.FieldLogger) func(r Model) {
	return func(r Model) {
		TimeoutRegistry().Cancel(r.Id())
	}
}

func refreshTimeout(l logrus.FieldLogger, span opentracing.Span) func(r Model) {
	return func(r Model) {
		to := r.Timeout()
		if to > -1 {
			ns := r.TimeoutState()
			TimeoutRegistry().Schedule(r.Id(), TryForceHitReactor(l, span)(r.Id(), ns), time.Duration(to)*time.Second)
		}
	}
}

func SetEventState(reactorId uint32, newState byte) (Model, error) {
	return GetRegistry().Update(reactorId, setEventState(newState))
}

func TryForceHitReactor(l logrus.FieldLogger, span opentracing.Span) func(reactorId uint32, newState int8) func() {
	return func(reactorId uint32, newState int8) func() {
		return func() {
			r, err := GetRegistry().Update(reactorId, setState(newState), shouldCollect(true))
			if err != nil {
				return
			}
			clearTimeout(l)(r)
			refreshTimeout(l, span)(r)
			searchItem(l)(r)
			trigger(l, span)(r, 0)
		}
	}
}

func searchItem(_ logrus.FieldLogger) func(r Model) {
	return func(r Model) {
		//TODO do not know if this is "necessary" at this point
	}
}

func Touch(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB) func(id uint32, characterId uint32) {
	return func(id uint32, characterId uint32) {
		IfPresent(id, performScriptAction(l, span, db)(characterId, script.InvokeTouch))
	}
}

func Release(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB) func(id uint32, characterId uint32) {
	return func(id uint32, characterId uint32) {
		IfPresent(id, performScriptAction(l, span, db)(characterId, script.InvokeRelease))
	}
}

func IfPresent(id uint32, op model.Operator[Model]) {
	model.IfPresent[Model](ByIdProvider(id), op)
}

type CharacterAction func(uint32, script.Action) model.Operator[Model]

func performScriptAction(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB) CharacterAction {
	return func(characterId uint32, action script.Action) model.Operator[Model] {
		return func(m Model) error {
			c := script.Context{WorldId: m.WorldId(), ChannelId: m.ChannelId(), MapId: m.MapId(), CharacterId: characterId, ReactorClassification: m.Classification(), ReactorId: m.Id()}
			s, err := registry2.GetRegistry().GetScript(m.Classification())
			if err != nil {
				return err
			}
			action(l, span, db)(c)(*s)
			return nil
		}
	}
}

func IsRecentHitFromAttack(l logrus.FieldLogger) func(id uint32) bool {
	return func(id uint32) bool {
		r, err := GetById(id)
		if err != nil {
			l.WithError(err).Errorf("Unable to locate reactor by id %d, assuming no recent hit.", id)
		}
		return r.AttackHit()
	}
}
