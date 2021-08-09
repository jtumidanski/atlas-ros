package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New6109006() script.Script {
	return generic.NewReactor(6109006, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		generic.EventBlueMessage("RELIC_OF_MASTERY_AWARDED")
		next := event.GetProperty(l)(e.Id(), "glpq5") + 1
		event.SetProperty(l)(e.Id(), "glpq5", next)
		if next == 5 {
			generic.EventBlueMessage("ANTELLION_NEXT")
			generic.ShowClearEffectInMapWithMapObject(610030500, "5pt", 2)
			generic.GiveEventParticipantsStageReward(5)
		}
	}))
}
