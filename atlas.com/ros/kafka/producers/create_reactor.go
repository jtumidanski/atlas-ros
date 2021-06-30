package producers

import "github.com/sirupsen/logrus"

type createReactorCommand struct {
	WorldId   byte   `json:"world_id"`
	ChannelId byte   `json:"channel_id"`
	MapId     uint32 `json:"map_id"`
	ReactorId uint32 `json:"reactor_id"`
	Name      string `json:"name"`
	State     byte   `json:"state"`
	X         int16  `json:"x"`
	Y         int16  `json:"y"`
	Delay     uint32 `json:"delay"`
	Direction byte   `json:"direction"`
}

func CreateReactor(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state byte, x int16, y int16, delay uint32, direction byte) {
	producer := ProduceEvent(l, "TOPIC_CREATE_REACTOR_COMMAND")
	return func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state byte, x int16, y int16, delay uint32, direction byte) {
		command := &createReactorCommand{
			WorldId:   worldId,
			ChannelId: channelId,
			MapId:     mapId,
			ReactorId: reactorId,
			Name:      name,
			State:     state,
			X:         x,
			Y:         y,
			Delay:     delay,
			Direction: direction,
		}
		producer(CreateKey(int(reactorId)), command)
	}
}
