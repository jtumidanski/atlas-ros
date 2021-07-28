package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewHuntersAltar2() script.Script {
	return generic.NewReactor(reactor.HuntersAltar2, generic.SetAct(generic.NoOp), generic.SetHit(HuntersAltar2Hit))
}

func HuntersAltar2Hit(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	//rm.hitMonsterWithReactor(6090001, 4)
	//rm.getReactor().setEventState(Math.floor(Math.random() * 3).byteValue())
}
