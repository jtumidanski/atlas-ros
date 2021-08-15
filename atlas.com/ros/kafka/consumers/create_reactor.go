package consumers

import (
	"atlas-ros/kafka/handler"
	"atlas-ros/reactor"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func EmptyCreateReactorCommand() handler.EmptyEventCreator {
	return func() interface{} {
		return &createReactorCommand{}
	}
}

func HandleCreateReactorCommand(_ *gorm.DB) handler.EventHandler {
	return func(l logrus.FieldLogger, e interface{}) {
		if command, ok := e.(*createReactorCommand); ok {
			_, err := reactor.Create(l)(command.WorldId, command.ChannelId, command.MapId, command.Classification, command.Name, command.State, command.X, command.Y, command.Delay, command.Direction)
			if err != nil {
				l.WithError(err).Errorf("Unable to create reactor %d in map %d by command.", command.Classification, command.MapId)
			}
		} else {
			l.Errorf("Unable to cast command provided to handler")
		}
	}
}
