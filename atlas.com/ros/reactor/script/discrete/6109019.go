package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New6109019() script.Script {
	return generic.NewReactor(6109019, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		generic.EventBlueMessage("THIEF_SIGIL_ACTIVATED")
		next := event.GetProperty(l)(e.Id(), "glpq4") + 1
		event.SetProperty(l)(e.Id(), "glpq4", next)
		if next == 5 {
			generic.EventBlueMessage("ANTELLION_NEXT")
			generic.ShowClearEffectInMapWithMapObject(610030400, "4pt", 2)
			generic.GiveEventParticipantsStageReward(4)
		}
	}))
}
