package discrete

import (
	"atlas-ros/character"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Warp2202002() script.ActFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		mapId := uint32(922000009)
		if character.QuestActive(l)(c.CharacterId, 3238) {
			mapId = 922000020
		}
		generic.WarpById(mapId, 0)(l, db, c)
	}
}
