package _map

import (
	"atlas-ros/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type createReactorCommand struct {
	WorldId        byte   `json:"world_id"`
	ChannelId      byte   `json:"channel_id"`
	MapId          uint32 `json:"map_id"`
	Classification uint32 `json:"classification"`
	Name           string `json:"name"`
	State          int8   `json:"state"`
	X              int16  `json:"x"`
	Y              int16  `json:"y"`
	Delay          uint32 `json:"delay"`
	Direction      byte   `json:"direction"`
}

func emitCreateReactor(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, classification uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_CREATE_REACTOR_COMMAND")
	return func(worldId byte, channelId byte, mapId uint32, classification uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) {
		command := &createReactorCommand{
			WorldId:        worldId,
			ChannelId:      channelId,
			MapId:          mapId,
			Classification: classification,
			Name:           name,
			State:          state,
			X:              x,
			Y:              y,
			Delay:          delay,
			Direction:      direction,
		}
		producer(kafka.CreateKey(int(classification)), command)
	}
}
