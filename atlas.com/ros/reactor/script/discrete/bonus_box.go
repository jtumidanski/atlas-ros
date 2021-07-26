package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
)

func NewBonusBox() script.Script {
	return generic.NewReactor(reactor.BonusBox, generic.SetAct(BonusBoxAct))
}

func BonusBoxAct(l logrus.FieldLogger, c script.Context) {
	//rm.dropItems(true, 1, 100, 400, 15)
	//
	//EventInstanceManager eim = rm.getEventInstance()
	//if (eim.getProperty("statusStgBonus") != "1") {
	//	rm.spawnNpc(2013002, new Point(46, 840))
	//	eim.setProperty("statusStgBonus", "1")
	//}
}
