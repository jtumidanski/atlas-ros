package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Hit2612005() script.HitFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		r, err := reactor.GetById(c.ReactorId)
		if err != nil {
			return
		}
		if r.State() == 4 {
			generic.Drop(false, 0, 0, 0, 0)(l, db, c)
		}
	}
}
