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

func NewRecord() script.Script {
	return generic.NewReactor(reactor.Record, generic.SetAct(RecordAct))
}

func RecordAct(l logrus.FieldLogger, _ opentracing.Span, _ *gorm.DB, c script.Context) {
	if !event.ParticipatingInEvent(l)(c.CharacterId) {
		return
	}

	e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
	if err != nil {
		return
	}
	event.SetStringProperty(l)(e.Id(), "statusStg3", "0")
}
