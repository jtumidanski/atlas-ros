package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Act2511000() script.ActFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}
		next := event.GetProperty(l)(e.Id(), "openedBoxes") + 1
		event.SetProperty(l)(e.Id(), "openedBoxes", next)
		generic.SpawnMonsters(9300109, 3)(l, db, c)
		generic.SpawnMonsters(9300110, 5)(l, db, c)
	}
}
