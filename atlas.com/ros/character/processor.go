package character

import (
	"atlas-ros/kafka/producers"
	"atlas-ros/portal"
	"github.com/sirupsen/logrus"
)

func GetDropRate(_ logrus.FieldLogger) func(characterId uint32) float64 {
	return func(characterId uint32) float64 {
		//TODO
		return 1
	}
}

func NeedsQuestItem(_ logrus.FieldLogger) func(characterId uint32, itemId uint32, questId int32) bool {
	return func(characterId uint32, itemId uint32, questId int32) bool {
		//TODO
		return false
	}
}

func WarpToPortal(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) {
		producers.ChangeMap(l)(worldId, channelId, characterId, mapId, p())
	}
}

func WarpById(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
		WarpToPortal(l)(worldId, channelId, characterId, mapId, portal.FixedPortalIdProvider(portalId))
	}
}

func WarpByName(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
		WarpToPortal(l)(worldId, channelId, characterId, mapId, portal.ByNamePortalIdProvider(l)(mapId, portalName))
	}
}
