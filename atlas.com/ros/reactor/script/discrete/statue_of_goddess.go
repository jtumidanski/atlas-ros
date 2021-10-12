package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func NewStatueOfGoddess() script.Script {
	return generic.NewReactor(reactor.StatueOfGoddess, generic.SetAct(StatueOfGoddessAct))
}

func StatueOfGoddessAct(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
	generic.SpawnNPC(2013002)(l, span, db, c)
	if !event.ParticipatingInEvent(l)(c.CharacterId) {
		return
	}

	e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
	if err != nil {
		return
	}
	event.ClearPartyQuest(l)(e.Id())
	event.SetStringProperty(l)(e.Id(), "statusStg8", "1")
	event.GiveParticipantsExperience(l)(e.Id(), 3500)
	generic.ShowClearEffectWithGate(true)(l, db, c)
	event.StartTimer(l)(e.Id(), 5 * time.Minute)
}
