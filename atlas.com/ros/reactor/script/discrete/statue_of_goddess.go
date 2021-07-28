package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewStatueOfGoddess() script.Script {
	return generic.NewReactor(reactor.StatueOfGoddess, generic.SetAct(StatueOfGoddessAct))
}

func StatueOfGoddessAct(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	//rm.spawnNpc(2013002)
	//rm.getEventInstance().clearPQ()
	//
	//rm.getEventInstance().setProperty("statusStg8", "1")
	//
	//EventInstanceManager eim = rm.getEventInstance()
	//eim.giveEventPlayersExp(3500)
	//eim.showClearEffect(true)
	//
	//rm.getEventInstance().startEventTimer(5 * 60000) //bonus time
}
