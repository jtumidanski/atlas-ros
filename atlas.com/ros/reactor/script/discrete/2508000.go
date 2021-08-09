package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Warp2508000() script.ActFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		if c.MapId/100%100 != 38 {
			generic.WarpRandom(c.MapId+100)(l, db, c)
		}
	}
}
