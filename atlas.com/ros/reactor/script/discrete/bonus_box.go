package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewBonusBox() script.Script {
	return generic.NewReactor(reactor.BonusBox, generic.SetAct(BonusBoxAct))
}

func BonusBoxAct(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
	if !event.ParticipatingInEvent(l)(c.CharacterId) {
		return
	}
	e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
	if err != nil {
		return
	}

	generic.Drop(true, 1, 100, 400, 15)(l, span, db, c)
	if event.GetStringProperty(l)(e.Id(), "statusStgBonus") != "1" {
		generic.SpawnNPCAt(2013002, 46, 840)(l, span, db, c)
		event.SetStringProperty(l)(e.Id(), "statusStgBonus", "1")
	}
}
