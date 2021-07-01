package reactor

import (
	"atlas-ros/kafka/producers"
	"atlas-ros/reactor/statistics"
	"github.com/sirupsen/logrus"
)

func Create(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state byte, x int16, y int16, delay uint32, direction byte) (*Model, error) {
	return func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state byte, x int16, y int16, delay uint32, direction byte) (*Model, error) {
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
