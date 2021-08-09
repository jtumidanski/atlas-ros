package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	rand2 "math/rand"
)

func New6702004() script.Script {
	return generic.NewReactor(6702004, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		rand := rand2.Intn(3) + 1
		for i := 0; i < rand; i++ {
			generic.Spray(true, 1, 30, 60, 15)(l, db, c)
		}
	}))
}
