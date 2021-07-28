package reactor

import (
	"atlas-ros/kafka/producers"
	"atlas-ros/reactor/script"
	registry2 "atlas-ros/reactor/script/registry"
	"atlas-ros/reactor/statistics"
	"github.com/sirupsen/logrus"
	"time"
)

func Create(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) (*Model, error) {
	return func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) (*Model, error) {
		s, err := statistics.GetCache().GetFile(reactorId)
		if err != nil {
			l.WithError(err).Errorf("Unable to get reactor statistics from WZ")
			return nil, err
		}
		m := GetRegistry().Create(worldId, channelId, mapId, reactorId, name, state, x, y, delay, direction, *s)
		producers.Created(l)(worldId, channelId, mapId, m.UniqueId())
		return &m, nil
	}
}

func GetInMap(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) []Model {
	return func(worldId byte, channelId byte, mapId uint32) []Model {
		return GetRegistry().GetInMap(worldId, channelId, mapId)
	}
}

func Get(_ logrus.FieldLogger) func(uniqueId uint32) (*Model, error) {
	return func(uniqueId uint32) (*Model, error) {
		return GetRegistry().Get(uniqueId)
	}
}

func Hit(l logrus.FieldLogger) func(uniqueId uint32, characterId uint32, stance uint16, skillId uint32) error {
	return func(uniqueId uint32, characterId uint32, stance uint16, skillId uint32) error {
		r, err := Get(l)(uniqueId)
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
					r, err = GetRegistry().Update(uniqueId, advanceState(i))
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
						r, err = GetRegistry().Update(uniqueId, shouldCollect(true))
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
			r, err = GetRegistry().Update(uniqueId, incrementState(), shouldCollect(true))
			if err != nil {
				return err
			}
			trigger(l)(r, stance)
			if r.Id() != 9980000 && r.Id() != 9980001 {
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
		producers.Triggered(l)(r.WorldId(), r.ChannelId(), r.MapId(), r.UniqueId(), stance)
	}
}

func destroy(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		GetRegistry().Destroy(r.UniqueId())
		clearTimeout(l)(r)
		producers.Destroyed(l)(r.WorldId(), r.ChannelId(), r.MapId(), r.UniqueId())
		go respawn(l)(r)
	}
}

func respawn(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		time.Sleep(time.Duration(r.Delay()) * time.Second)
		_, err := Create(l)(r.WorldId(), r.ChannelId(), r.MapId(), r.Id(), r.Name(), 0, r.X(), r.Y(), r.Delay(), r.FacingDirection())
		if err != nil {
			l.WithError(err).Errorf("Unable to respawn reactor %d at location %d,%d in map %d.", r.Id(), r.X(), r.Y(), r.MapId())
		}
	}
}

func clearTimeout(_ logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		TimeoutRegistry().Cancel(r.UniqueId())
	}
}

func refreshTimeout(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {
		to := r.Timeout()
		if to > -1 {
			ns := r.TimeoutState()
			TimeoutRegistry().Schedule(r.UniqueId(), tryForceHitReactor(l)(r.UniqueId(), ns), time.Duration(to)*time.Second)
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

func Touch(l logrus.FieldLogger) func(uniqueId uint32, characterId uint32) error {
	return func(uniqueId uint32, characterId uint32) error {
		return For(uniqueId, performScriptAction(l)(characterId, script.InvokeTouch))
	}
}

func Release(l logrus.FieldLogger) func(uniqueId uint32, characterId uint32) error {
	return func(uniqueId uint32, characterId uint32) error {
		return For(uniqueId, performScriptAction(l)(characterId, script.InvokeRelease))
	}
}

// For applies a Operator function on a reactor, given its id
func For(uniqueId uint32, rf Operator) error {
	r, err := GetRegistry().Get(uniqueId)
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
			c := script.Context{WorldId: m.WorldId(), ChannelId: m.ChannelId(), MapId: m.MapId(), CharacterId: characterId, ReactorId: m.Id(), ReactorUniqueId: m.UniqueId()}
			s, err := registry2.GetRegistry().GetScript(m.Id())
			if err != nil {
				return err
			}
			action(l)(c)(*s)
			return nil
		}
	}
}
