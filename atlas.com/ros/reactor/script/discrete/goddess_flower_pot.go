package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewGoddessFlowerPot() script.Script {
	return generic.NewReactor(reactor.GoddessFlowerPot, generic.SetAct(GoddessFlowerPotAct))
}

func GoddessFlowerPotAct(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	//if (rm.getMap().getSummonState()) {
	//	int count = rm.getEventInstance().getIntProperty("statusStg7_c")
	//
	//	if (count < 7) {
	//		int nextCount = (count + 1)
	//
	//		rm.spawnMonster(Math.random() >= 0.6 ? 9300049 : 9300048)
	//		rm.getEventInstance().setProperty("statusStg7_c", nextCount)
	//	} else {
	//		rm.spawnMonster(9300049)
	//	}
	//}
}
