package reactor

import (
	"atlas-ros/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type reactorStatusEvent struct {
	WorldId   byte   `json:"world_id"`
	ChannelId byte   `json:"channel_id"`
	MapId     uint32 `json:"map_id"`
	Id        uint32 `json:"id"`
	Status    string `json:"status"`
	Stance    uint16 `json:"stance"`
}

func emitCreated(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, id uint32) {
	f := emitStatusEvent(l, span)
	return func(worldId byte, channelId byte, mapId uint32, id uint32) {
		f(worldId, channelId, mapId, id, "CREATED", 0)
	}
}

func emitTriggered(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, id uint32, stance uint16) {
	f := emitStatusEvent(l, span)
	return func(worldId byte, channelId byte, mapId uint32, id uint32, stance uint16) {
		f(worldId, channelId, mapId, id, "TRIGGERED", stance)
	}
}

func emitDestroyed(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, id uint32) {
	f := emitStatusEvent(l, span)
	return func(worldId byte, channelId byte, mapId uint32, id uint32) {
		f(worldId, channelId, mapId, id, "DESTROYED", 0)
	}
}

func emitStatusEvent(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, id uint32, status string, stance uint16) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_REACTOR_STATUS_EVENT")
	return func(worldId byte, channelId byte, mapId uint32, id uint32, status string, stance uint16) {
		e := &reactorStatusEvent{WorldId: worldId, ChannelId: channelId, MapId: mapId, Id: id, Status: status, Stance: stance}
		producer(kafka.CreateKey(int(id)), e)
	}
}
