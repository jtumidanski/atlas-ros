package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Hit2612005() script.HitFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		r, err := reactor.GetById(c.ReactorId)
		if err != nil {
			return
		}
		if r.State() == 4 {
			generic.Drop(false, 0, 0, 0, 0)(l, span, db, c)
		}
	}
}
