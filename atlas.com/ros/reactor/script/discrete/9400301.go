package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New9400301() script.Script {
	return generic.NewReactor(9400301, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		event.InvokeFunction(l)(e.Id(), "snowmanSnackFake")
	}))
}
