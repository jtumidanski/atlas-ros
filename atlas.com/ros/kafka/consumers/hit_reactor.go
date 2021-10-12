package consumers

import (
	"atlas-ros/kafka/handler"
	"atlas-ros/reactor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type hitReactorCommand struct {
	WorldId     byte   `json:"world_id"`
	ChannelId   byte   `json:"channel_id"`
	MapId       uint32 `json:"map_id"`
	CharacterId uint32 `json:"character_id"`
	Id          uint32 `json:"id"`
	Stance      uint16 `json:"stance"`
	SkillId     uint32 `json:"skill_id"`
}

func EmptyHitReactorCommand() handler.EmptyEventCreator {
	return func() interface{} {
		return &hitReactorCommand{}
	}
}

func HandleHitReactorCommand(db *gorm.DB) handler.EventHandler {
	return func(l logrus.FieldLogger, span opentracing.Span, e interface{}) {
		if command, ok := e.(*hitReactorCommand); ok {
			err := reactor.Hit(l, span, db)(command.Id, command.CharacterId, command.Stance, command.SkillId)
			if err != nil {
				l.WithError(err).Errorf("Unable to hit reactor %d in map %d by command.", command.Id, command.MapId)
			}
		} else {
			l.Errorf("Unable to cast command provided to handler")
		}
	}
}
