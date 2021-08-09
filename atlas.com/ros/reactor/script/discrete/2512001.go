package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Act2512001() script.ActFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		next := event.GetProperty(l)(e.Id(), "openedChests")
		event.SetProperty(l)(e.Id(), "openedChests", next)
		generic.Spray(true, 1, 50, 100, 15)(l, db, c)
	}
}
