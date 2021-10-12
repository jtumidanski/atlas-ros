package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Act2511000() script.ActFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}
		next := event.GetProperty(l)(e.Id(), "openedBoxes") + 1
		event.SetProperty(l)(e.Id(), "openedBoxes", next)
		generic.SpawnMonsters(9300109, 3)(l, span, db, c)
		generic.SpawnMonsters(9300110, 5)(l, span, db, c)
	}
}
