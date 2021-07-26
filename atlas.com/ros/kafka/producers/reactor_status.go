package producers

import "github.com/sirupsen/logrus"

type reactorStatusEvent struct {
	WorldId   byte   `json:"world_id"`
	ChannelId byte   `json:"channel_id"`
	MapId     uint32 `json:"map_id"`
	UniqueId  uint32 `json:"unique_id"`
	Status    string `json:"status"`
	Stance    uint16 `json:"stance"`
}

func Created(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32) {
	f := StatusEvent(l)
	return func(worldId byte, channelId byte, mapId uint32, uniqueId uint32) {
		f(worldId, channelId, mapId, uniqueId, "CREATED", 0)
	}
}

func Triggered(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, stance uint16) {
	f := StatusEvent(l)
	return func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, stance uint16) {
		f(worldId, channelId, mapId, uniqueId, "TRIGGERED", stance)
	}
}

func Destroyed(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32) {
	f := StatusEvent(l)
	return func(worldId byte, channelId byte, mapId uint32, uniqueId uint32) {
		f(worldId, channelId, mapId, uniqueId, "DESTROYED", 0)
	}
}

func StatusEvent(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, status string, stance uint16) {
	producer := ProduceEvent(l, "TOPIC_REACTOR_STATUS_EVENT")
	return func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, status string, stance uint16) {
		e := &reactorStatusEvent{WorldId: worldId, ChannelId: channelId, MapId: mapId, UniqueId: uniqueId, Status: status, Stance: stance}
		producer(CreateKey(int(uniqueId)), e)
	}
}
