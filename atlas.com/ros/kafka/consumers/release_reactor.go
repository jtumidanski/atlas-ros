package consumers

import (
	"atlas-ros/kafka/handler"
	"atlas-ros/reactor"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type releaseReactorCommand struct {
	WorldId     byte   `json:"world_id"`
	ChannelId   byte   `json:"channel_id"`
	MapId       uint32 `json:"map_id"`
	CharacterId uint32 `json:"character_id"`
	Id          uint32 `json:"id"`
}

func EmptyReleaseReactorCommand() handler.EmptyEventCreator {
	return func() interface{} {
		return &releaseReactorCommand{}
	}
}

func HandleReleaseReactorCommand(db *gorm.DB) handler.EventHandler {
	return func(l logrus.FieldLogger, e interface{}) {
		if command, ok := e.(*releaseReactorCommand); ok {
			err := reactor.Release(l)(command.Id, command.CharacterId)
			if err != nil {
				l.WithError(err).Errorf("Unable to release reactor %d in map %d by command.", command.Id, command.MapId)
			}
		} else {
			l.Errorf("Unable to cast command provided to handler")
		}
	}
}
