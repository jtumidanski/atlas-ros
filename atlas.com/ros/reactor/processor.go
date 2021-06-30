package reactor

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Create(l logrus.FieldLogger, db *gorm.DB) func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state byte, x int16, y int16, delay uint32, direction byte) (*Model, error) {
	return func(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state byte, x int16, y int16, delay uint32, direction byte) (*Model, error) {
		return nil, nil
	}
}
