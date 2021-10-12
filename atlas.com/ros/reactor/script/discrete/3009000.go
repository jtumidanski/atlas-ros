package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New3009000() script.Script {
	return generic.NewReactor(3009000, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		r, err := reactor.GetById(c.ReactorId)
		if err != nil {
			return
		}
		if r.State() == 4 {
			generic.ShowClearEffect()(l, span, db, c)
		}
	}))
}
