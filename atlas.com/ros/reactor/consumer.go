package reactor

import (
	"atlas-ros/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	consumerNameCreate  = "create_reactor_command"
	consumerNameHit     = "hit_reactor_command"
	consumerNameTouch   = "touch_reactor_command"
	consumerNameRelease = "release_reactor_command"
	topicTokenCreate    = "TOPIC_CREATE_REACTOR_COMMAND"
	topicTokenHit       = "TOPIC_HIT_REACTOR_COMMAND"
	topicTokenTouch     = "TOPIC_TOUCH_REACTOR_COMMAND"
	topicTokenRelease   = "TOPIC_RELEASE_REACTOR_COMMAND"
)

func CreateConsumer(db *gorm.DB) func(groupId string) kafka.ConsumerConfig {
	return func(groupId string) kafka.ConsumerConfig {
		return kafka.NewConsumerConfig[createReactorCommand](consumerNameCreate, topicTokenCreate, groupId, handleCreate(db))
	}
}

func handleCreate(_ *gorm.DB) kafka.HandlerFunc[createReactorCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command createReactorCommand) {
		_, err := Create(l, span)(command.WorldId, command.ChannelId, command.MapId, command.Classification, command.Name, command.State, command.X, command.Y, command.Delay, command.Direction)
		if err != nil {
			l.WithError(err).Errorf("Unable to create reactor %d in map %d by command.", command.Classification, command.MapId)
		}
	}
}

func HitConsumer(db *gorm.DB) func(groupId string) kafka.ConsumerConfig {
	return func(groupId string) kafka.ConsumerConfig {
		return kafka.NewConsumerConfig[hitReactorCommand](consumerNameHit, topicTokenHit, groupId, handleHit(db))
	}
}

type hitReactorCommand struct {
	WorldId     byte   `json:"world_id"`
	ChannelId   byte   `json:"channel_id"`
	MapId       uint32 `json:"map_id"`
	CharacterId uint32 `json:"character_id"`
	Id          uint32 `json:"id"`
	Stance      uint16 `json:"stance"`
	SkillId     uint32 `json:"skill_id"`
}

func handleHit(db *gorm.DB) kafka.HandlerFunc[hitReactorCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command hitReactorCommand) {
		err := Hit(l, span, db)(command.Id, command.CharacterId, command.Stance, command.SkillId)
		if err != nil {
			l.WithError(err).Errorf("Unable to hit reactor %d in map %d by command.", command.Id, command.MapId)
		}
	}
}

func ReleaseConsumer(db *gorm.DB) func(groupId string) kafka.ConsumerConfig {
	return func(groupId string) kafka.ConsumerConfig {
		return kafka.NewConsumerConfig[releaseReactorCommand](consumerNameRelease, topicTokenRelease, groupId, handleRelease(db))
	}
}

type releaseReactorCommand struct {
	WorldId     byte   `json:"world_id"`
	ChannelId   byte   `json:"channel_id"`
	MapId       uint32 `json:"map_id"`
	CharacterId uint32 `json:"character_id"`
	Id          uint32 `json:"id"`
}

func handleRelease(db *gorm.DB) kafka.HandlerFunc[releaseReactorCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command releaseReactorCommand) {
		Release(l, span, db)(command.Id, command.CharacterId)
	}
}

func TouchConsumer(db *gorm.DB) func(groupId string) kafka.ConsumerConfig {
	return func(groupId string) kafka.ConsumerConfig {
		return kafka.NewConsumerConfig[touchReactorCommand](consumerNameTouch, topicTokenTouch, groupId, handleTouch(db))
	}
}

type touchReactorCommand struct {
	WorldId     byte   `json:"world_id"`
	ChannelId   byte   `json:"channel_id"`
	MapId       uint32 `json:"map_id"`
	CharacterId uint32 `json:"character_id"`
	Id          uint32 `json:"id"`
}

func handleTouch(db *gorm.DB) kafka.HandlerFunc[touchReactorCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command touchReactorCommand) {
		Touch(l, span, db)(command.Id, command.CharacterId)
	}
}
