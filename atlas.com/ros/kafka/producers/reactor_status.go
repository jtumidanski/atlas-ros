package producers

import "github.com/sirupsen/logrus"

type reactorStatusEvent struct {
	WorldId   byte   `json:"world_id"`
	ChannelId byte   `json:"channel_id"`
	MapId     uint32 `json:"map_id"`
	UniqueId  uint32 `json:"unique_id"`
	Status    string `json:"status"`
}

func Created(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32) {
	f := StatusEvent(l)
	return func(worldId byte, channelId byte, mapId uint32, uniqueId uint32) {
		f(worldId, channelId, mapId, uniqueId, "CREATED")
	}
}

func StatusEvent(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, status string) {
	producer := ProduceEvent(l, "TOPIC_REACTOR_STATUS_EVENT")
	return func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, status string) {
		e := &reactorStatusEvent{WorldId: worldId, ChannelId: channelId, MapId: mapId, UniqueId: uniqueId, Status: status}
		producer(CreateKey(int(uniqueId)), e)
	}
}
