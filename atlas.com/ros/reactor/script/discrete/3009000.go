package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New3009000() script.Script {
	return generic.NewReactor(3009000, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		r, err := reactor.GetById(l)(c.ReactorId)
		if err != nil {
			return
		}
		if r.State() == 4 {
			generic.ShowClearEffect()(l, db, c)
		}
	}))
}
