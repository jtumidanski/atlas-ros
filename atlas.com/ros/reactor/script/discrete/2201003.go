package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Warp2201003() script.ActFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if c.MapId == 922010900 {
			generic.MapPinkMessage("ALISHAR_SUMMONED")(l, span, db, c)
			generic.SpawnMonsterAt(9300012, 941, 184)(l, span, db, c)
		} else if c.MapId == 922010700 {
			generic.MapPinkMessage("ROMBARD_SUMMONED")(l, span, db, c)
			generic.SpawnMonsterAt(9300010, 1, -211)(l, span, db, c)
		}
	}
}
