package character

import (
	"atlas-ros/kafka/producers"
	"atlas-ros/portal"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetDropRate(_ logrus.FieldLogger) func(characterId uint32) float64 {
	return func(characterId uint32) float64 {
		//TODO AT-3
		return 1
	}
}

func NeedsQuestItem(_ logrus.FieldLogger) func(characterId uint32, itemId uint32, questId int32) bool {
	return func(characterId uint32, itemId uint32, questId int32) bool {
		//TODO AT-4
		return false
	}
}

func WarpToPortal(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) {
		producers.ChangeMap(l, span)(worldId, channelId, characterId, mapId, p())
	}
}

func WarpById(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, portal.FixedPortalIdProvider(portalId))
	}
}

func WarpByName(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, portal.ByNamePortalIdProvider(l, span)(mapId, portalName))
	}
}

func WarpRandom(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, portal.RandomPortalIdProvider(l, span)(mapId))
	}
}

func QuestActive(_ logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		// TODO AT-5
		return false
	}
}

func QuestStarted(_ logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		// TODO AT-5
		return false
	}
}

func SetQuestProgress(_ logrus.FieldLogger) func(characterId uint32, questId uint32, infoNumber uint32, value int32) {
	return func(characterId uint32, questId uint32, infoNumber uint32, value int32) {
		// TODO AT-6
	}
}

func SetQuestProgressString(_ logrus.FieldLogger) func(characterId uint32, questId uint32, infoNumber uint32, value string) {
	return func(characterId uint32, questId uint32, infoNumber uint32, value string) {
		// TODO AT-6
	}
}

func SendNotice(l logrus.FieldLogger) func(characterId uint32, noticeType string, message string) {
	return func(characterId uint32, noticeType string, message string) {
		//TODO AT-7
	}
}
