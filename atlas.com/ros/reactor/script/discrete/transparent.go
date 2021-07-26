package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
)

func NewTransparent() script.Script {
	return generic.NewReactor(reactor.Transparent, generic.SetAct(TransparentAct))
}

func TransparentAct(l logrus.FieldLogger, c script.Context) {
	//if (rm.getEventInstance().getIntProperty("statusStg2") == -1) {
	//	int rnd = Math.max(Math.floor(Math.random() * 14), 4).intValue()
	//
	//	rm.getEventInstance().setProperty("statusStg2", "" + rnd)
	//	rm.getEventInstance().setProperty("statusStg2_c", "0")
	//}
	//
	//int limit = rm.getEventInstance().getIntProperty("statusStg2")
	//int count = rm.getEventInstance().getIntProperty("statusStg2_c")
	//if (count >= limit) {
	//	rm.dropItems()
	//
	//	EventInstanceManager eim = rm.getEventInstance()
	//	eim.giveEventPlayersExp(3500)
	//
	//	eim.setProperty("statusStg2", "1")
	//	eim.showClearEffect(true)
	//} else {
	//	count++
	//	rm.getEventInstance().setProperty("statusStg2_c", count)
	//
	//	int nextHashed = (11 * (count)) % 14
	//
	//	Point nextPos = rm.getMap().getReactorById(2001002 + nextHashed).position()
	//	rm.spawnMonster(9300040, 1, nextPos)
	//}
}
