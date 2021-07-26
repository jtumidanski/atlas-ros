package reactor

import (
	"atlas-ros/kafka/producers"
	"atlas-ros/reactor/script"
	registry2 "atlas-ros/reactor/script/registry"
	"atlas-ros/reactor/statistics"
	"github.com/sirupsen/logrus"
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

func Hit(l logrus.FieldLogger) func(uniqueId uint32, characterId uint32, characterPosition uint32, stance uint16, skillId uint32) error {
	return func(uniqueId uint32, characterId uint32, characterPosition uint32, stance uint16, skillId uint32) error {
		r, err := Get(l)(uniqueId)
		if err != nil {
			return err
		}
		if !r.Active() {
			return nil
		}
		clearTimeout(l)(r)
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
							}
							trigger(l)(r, stance)
						} else {
							trigger(l)(r, stance)
						}
						act(l)(characterId, r)
					} else {
						trigger(l)(r, stance)
						if r.State() == r.NextState(i) {
							act(l)(characterId, r)
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
				act(l)(characterId, r)
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
		//trigger respawn
		producers.Destroyed(l)(r.WorldId(), r.ChannelId(), r.MapId(), r.UniqueId())
	}
}

func act(l logrus.FieldLogger) func(characterId uint32, r *Model) {
	return func(characterId uint32, r *Model) {
		s, err := registry2.GetRegistry().GetScript(r.Id())
		if err != nil {
			l.WithError(err).Errorf("Unable to locate script for reactor %d.", r.Id())
			return
		}
		(*s).Act(l, script.Context{WorldId: r.WorldId(), ChannelId: r.ChannelId(), MapId: r.MapId(), CharacterId: characterId, ReactorId: r.Id(), ReactorUniqueId: r.UniqueId()})
	}
}

func clearTimeout(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {

	}
}

func refreshTimeout(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {

	}
}

func searchItem(l logrus.FieldLogger) func(r *Model) {
	return func(r *Model) {

	}
}

func Touch(l logrus.FieldLogger) func(uniqueId uint32) error {
	return func(uniqueId uint32) error {
		return nil
	}
}

func Release(l logrus.FieldLogger) func(uniqueId uint32) error {
	return func(uniqueId uint32) error {
		return nil
	}
}
