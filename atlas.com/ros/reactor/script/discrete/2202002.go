package discrete

import (
	"atlas-ros/character"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Warp2202002() script.ActFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		mapId := uint32(922000009)
		if character.QuestActive(l)(c.CharacterId, 3238) {
			mapId = 922000020
		}
		generic.WarpById(mapId, 0)(l, span, db, c)
	}
}
