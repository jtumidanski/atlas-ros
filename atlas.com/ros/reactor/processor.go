package reactor

import (
	"atlas-ros/kafka/producers"
	"atlas-ros/reactor/script"
	registry2 "atlas-ros/reactor/script/registry"
	"atlas-ros/reactor/statistics"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func Create(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, classification uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) (*Model, error) {
	return func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) (*Model, error) {
		s, err := statistics.GetCache().GetFile(reactorId)
		if err != nil {
			l.WithError(err).Errorf("Unable to get reactor statistics from WZ")
			return nil, err
		}
		m := GetRegistry().Create(worldId, channelId, mapId, reactorId, name, state, x, y, delay, direction, *s)
		producers.Created(l)(worldId, channelId, mapId, m.Id())
		return &m, nil
	}
}

func GetInMap(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) []Model {
	return func(worldId byte, channelId byte, mapId uint32) []Model {
		return GetRegistry().GetInMap(worldId, channelId, mapId)
	}
}

type IdProvider func() uint32

func Get(provider IdProvider) (*Model, error) {
	id := provider()
	return GetRegistry().Get(id)
}

func FixedIdProvider(id uint32) IdProvider {
	return func() uint32 {
		return id
	}
}

func ByNameInMapProvider(worldId byte, channelId byte, mapId uint32, name string) IdProvider {
	return func() uint32 {
		for _, rs := range GetRegistry().GetInMap(worldId, channelId, mapId) {
			if name == rs.Name() {
				return rs.Id()
			}
		}
		return 0
	}
}

func ByClassificationInMapProvider(worldId byte, channelId byte, mapId uint32, classification uint32) IdProvider {
	return func() uint32 {
		for _, rs := range GetRegistry().GetInMap(worldId, channelId, mapId) {
			if classification == rs.Classification() {
				return rs.Id()
			}
		}
		return 0
	}
}

func GetById(id uint32) (*Model, error) {
	return Get(FixedIdProvider(id))
}

func GetByNameInMap(worldId byte, channelId byte, mapId uint32, name string) (*Model, error) {
	return Get(ByNameInMapProvider(worldId, channelId, mapId, name))
}

func GetByClassificationInMap(worldId byte, channelId byte, mapId uint32, classification uint32) (*Model, error) {
	return Get(ByClassificationInMapProvider(worldId, channelId, mapId, classification))
}

func Hit(l logrus.FieldLogger, db *gorm.DB) func(id uint32, characterId uint32, stance uint16, skillId uint32) error {
	return func(id uint32, characterId uint32, stance uint16, skillId uint32) error {
		r, err := GetById(id)
		if err != nil {
			return err
		}
		if !r.Active() {
			return nil
		}
		clearTimeout(l)(r)

		err = performScriptAction(l, db)(characterId, script.InvokeHit)(r)
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
								destroy(l)(r)
							} else {
								trigger(l)(r, stance)
							}
						} else {
							trigger(l)(r, stance)
						}
						err = performScriptAction(l, db)(characterId, script.InvokeAct)(r)
						if err != nil {
							return err
						}
					} else {
						trigger(l)(r, stance)
						if r.State() == r.NextState(i) {
							err = performScriptAction(l, db)(characterId, script.InvokeAct)(r)
							if err != nil {
								return err
							}
						}
						r, err = GetRegistry().Update(id, shouldCollect(true))
						if err != nil {
							return err
						}
						refreshTimeout(l)(r)
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
			trigger(l)(r, stance)
			if r.Classification() != 9980000 && r.Classification() != 9980001 {
				err = performScriptAction(l, db)(characterId, script.InvokeAct)(r)
				if err != nil {
					return err
				}
			}
			refreshTimeout(l)(r)
			if r.Type() == 100 {
				searchItem(l)(r)
			}
		}
		return nil
	}
}

func trigger(l logrus.FieldLogger) func(r *Model, stance uint16) {
	return func(r *Model, stance uint16) {
		producers.Triggered(l)(r.WorldId(), r.ChannelId(), r.MapId(), r.Id(), stance)
	}
}

func destroy(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		dr, err := GetRegistry().Destroy(r.Id())
		if err != nil {
			l.WithError(err).Errorf("Unable to destroy reactor %d.", r.Id())
			return
		}
		clearTimeout(l)(dr)
		producers.Destroyed(l)(dr.WorldId(), dr.ChannelId(), dr.MapId(), dr.Id())
		go respawn(l)(dr)
		//TODO need task to clean up destroyed reactors.
	}
}

func respawn(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		time.Sleep(time.Duration(r.Delay()) * time.Millisecond)
		_, err := Create(l)(r.WorldId(), r.ChannelId(), r.MapId(), r.Classification(), r.Name(), 0, r.X(), r.Y(), r.Delay(), r.FacingDirection())
		if err != nil {
			l.WithError(err).Errorf("Unable to respawn reactor %d at location %d,%d in map %d.", r.Classification(), r.X(), r.Y(), r.MapId())
		}
	}
}

func clearTimeout(_ logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		TimeoutRegistry().Cancel(r.Id())
	}
}

func refreshTimeout(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		to := r.Timeout()
		if to > -1 {
			ns := r.TimeoutState()
			TimeoutRegistry().Schedule(r.Id(), TryForceHitReactor(l)(r.Id(), ns), time.Duration(to)*time.Second)
		}
	}
}

func SetEventState(reactorId uint32, newState byte) (*Model, error) {
	return GetRegistry().Update(reactorId, setEventState(newState))
}

func TryForceHitReactor(l logrus.FieldLogger) func(reactorId uint32, newState int8) func() {
	return func(reactorId uint32, newState int8) func() {
		return func() {
			r, err := GetRegistry().Update(reactorId, setState(newState), shouldCollect(true))
			if err != nil {
				return
			}
			clearTimeout(l)(r)
			refreshTimeout(l)(r)
			searchItem(l)(r)
			trigger(l)(r, 0)
		}
	}
}

func searchItem(_ logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		//TODO do not know if this is "necessary" at this point
	}
}

func Touch(l logrus.FieldLogger, db *gorm.DB) func(id uint32, characterId uint32) error {
	return func(id uint32, characterId uint32) error {
		return For(id, performScriptAction(l, db)(characterId, script.InvokeTouch))
	}
}

func Release(l logrus.FieldLogger, db *gorm.DB) func(id uint32, characterId uint32) error {
	return func(id uint32, characterId uint32) error {
		return For(id, performScriptAction(l, db)(characterId, script.InvokeRelease))
	}
}

// For applies an Operator function on a reactor, given its id
func For(id uint32, rf Operator) error {
	r, err := GetRegistry().Get(id)
	if err != nil {
		return err
	}
	return rf(r)
}

// Operator a function which operators on a model
type Operator func(*Model) error

type CharacterAction func(uint32, script.Action) Operator

func performScriptAction(l logrus.FieldLogger, db *gorm.DB) CharacterAction {
	return func(characterId uint32, action script.Action) Operator {
		return func(m *Model) error {
			c := script.Context{WorldId: m.WorldId(), ChannelId: m.ChannelId(), MapId: m.MapId(), CharacterId: characterId, ReactorClassification: m.Classification(), ReactorId: m.Id()}
			s, err := registry2.GetRegistry().GetScript(m.Classification())
			if err != nil {
				return err
			}
			action(l, db)(c)(*s)
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
