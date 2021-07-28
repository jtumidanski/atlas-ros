package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewTransparent10() script.Script {
	return generic.NewReactor(reactor.Transparent10, generic.SetAct(Transparent10Act))
}

func Transparent10Act(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	//if (rm.getEventInstance().getIntProperty("statusStg2") == -1) {
	//	int rnd = Math.floor(Math.random() * 14).intValue()
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
