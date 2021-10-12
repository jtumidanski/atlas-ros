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

func NewGrassOfLife() script.Script {
	return generic.NewReactor(reactor.GrassOfLife, generic.SetAct(GrassOfLifeAct))
}

func GrassOfLifeAct(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
	generic.Drop(false, 0, 0, 0, 0)(l, span, db, c)
	if !event.ParticipatingInEvent(l)(c.CharacterId) {
		return
	}

	e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
	if err != nil {
		return
	}
	event.SetStringProperty(l)(e.Id(), "statusStg7", "1")
}
