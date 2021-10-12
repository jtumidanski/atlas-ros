package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New9208009() script.Script {
	return generic.NewReactor(9208009, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		event.SetStringProperty(l)(e.Id(), "canRevive", "1")
	}))
}
