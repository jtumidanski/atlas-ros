package consumers

import (
	"atlas-ros/kafka/handler"
	"atlas-ros/reactor"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type touchReactorCommand struct {
	WorldId   byte   `json:"world_id"`
	ChannelId byte   `json:"channel_id"`
	MapId     uint32 `json:"map_id"`
	ReactorId uint32 `json:"reactor_id"`
}

func EmptyTouchReactorCommand() handler.EmptyEventCreator {
	return func() interface{} {
		return &touchReactorCommand{}
	}
}

func HandleTouchReactorCommand(db *gorm.DB) handler.EventHandler {
	return func(l logrus.FieldLogger, e interface{}) {
		if command, ok := e.(*touchReactorCommand); ok {
			err := reactor.Touch(l)(command.WorldId, command.ChannelId, command.MapId, command.ReactorId)
			if err != nil {
				l.WithError(err).Errorf("Unable to touch reactor %d in map %d by command.", command.ReactorId, command.MapId)
			}
		} else {
			l.Errorf("Unable to cast command provided to handler")
		}
	}
}
