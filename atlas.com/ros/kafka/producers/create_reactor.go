package producers

import "github.com/sirupsen/logrus"

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

func CreateReactor(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, classification uint32, name string, state int8, x int16, y int16, delay uint32, direction byte) {
	producer := ProduceEvent(l, "TOPIC_CREATE_REACTOR_COMMAND")
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
		producer(CreateKey(int(classification)), command)
	}
}
