package reactor

import (
	"atlas-ros/kafka/producers"
	"atlas-ros/reactor/script"
	registry2 "atlas-ros/reactor/script/registry"
	"atlas-ros/reactor/statistics"
	"github.com/sirupsen/logrus"
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

func Get(_ logrus.FieldLogger) func(id uint32) (*Model, error) {
	return func(id uint32) (*Model, error) {
		return GetRegistry().Get(id)
	}
}

func Hit(l logrus.FieldLogger) func(id uint32, characterId uint32, stance uint16, skillId uint32) error {
	return func(id uint32, characterId uint32, stance uint16, skillId uint32) error {
		r, err := Get(l)(id)
		if err != nil {
			return err
		}
		if !r.Active() {
			return nil
		}
		clearTimeout(l)(r)

		err = performScriptAction(l)(characterId, script.InvokeHit)(r)
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
						err = performScriptAction(l)(characterId, script.InvokeAct)(r)
						if err != nil {
							return err
						}
					} else {
						trigger(l)(r, stance)
						if r.State() == r.NextState(i) {
							err = performScriptAction(l)(characterId, script.InvokeAct)(r)
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
				err = performScriptAction(l)(characterId, script.InvokeAct)(r)
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
		GetRegistry().Destroy(r.Id())
		clearTimeout(l)(r)
		producers.Destroyed(l)(r.WorldId(), r.ChannelId(), r.MapId(), r.Id())
		go respawn(l)(r)
	}
}

func respawn(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		time.Sleep(time.Duration(r.Delay()) * time.Second)
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
			TimeoutRegistry().Schedule(r.Id(), tryForceHitReactor(l)(r.Id(), ns), time.Duration(to)*time.Second)
		}
	}
}

func tryForceHitReactor(l logrus.FieldLogger) func(reactorId uint32, newState int8) func() {
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

func searchItem(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		//TODO do not know if this is "necessary" at this point
	}
}

func Touch(l logrus.FieldLogger) func(id uint32, characterId uint32) error {
	return func(id uint32, characterId uint32) error {
		return For(id, performScriptAction(l)(characterId, script.InvokeTouch))
	}
}

func Release(l logrus.FieldLogger) func(id uint32, characterId uint32) error {
	return func(id uint32, characterId uint32) error {
		return For(id, performScriptAction(l)(characterId, script.InvokeRelease))
	}
}

// For applies a Operator function on a reactor, given its id
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

func performScriptAction(l logrus.FieldLogger) CharacterAction {
	return func(characterId uint32, action script.Action) Operator {
		return func(m *Model) error {
			c := script.Context{WorldId: m.WorldId(), ChannelId: m.ChannelId(), MapId: m.MapId(), CharacterId: characterId, ReactorClassification: m.Classification(), ReactorId: m.Id()}
			s, err := registry2.GetRegistry().GetScript(m.Classification())
			if err != nil {
				return err
			}
			action(l)(c)(*s)
			return nil
		}
	}
}
