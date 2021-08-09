package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Warp2201003() script.ActFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		if c.MapId == 922010900 {
			generic.MapPinkMessage("ALISHAR_SUMMONED")(l, db, c)
			generic.SpawnMonsterAt(9300012, 941, 184)(l, db, c)
		} else if c.MapId == 922010700 {
			generic.MapPinkMessage("ROMBARD_SUMMONED")(l, db, c)
			generic.SpawnMonsterAt(9300010, 1, -211)(l, db, c)
		}
	}
}
