package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	rand2 "math/rand"
)

func New6702012() script.Script {
	return generic.NewReactor(6702012, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		rand := rand2.Intn(3) + 1
		for i := 0; i < rand; i++ {
			generic.Spray(true, 1, 30, 60, 15)(l, span, db, c)
		}
	}))
}
