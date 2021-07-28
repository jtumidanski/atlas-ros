package drop

import "github.com/sirupsen/logrus"

func Produce(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorClassification uint32, characterId uint32, meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) {
	return func(worldId byte, channelId byte, mapId uint32, reactorClassification uint32, characterId uint32, meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) {
		
	}
}
