package monster

import "github.com/sirupsen/logrus"

func Spawn(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, monsterId uint32, amount uint16, x int16, y int16) {
	return func(worldId byte, channelId byte, mapId uint32, monsterId uint32, amount uint16, x int16, y int16) {

	}
}
