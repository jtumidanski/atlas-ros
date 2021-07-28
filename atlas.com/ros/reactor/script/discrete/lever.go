package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewLever() script.Script {
	return generic.NewReactor(reactor.Lever, generic.SetAct(generic.NoOp), generic.SetHit(LeverHit))
}

func LeverHit(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	//MapleMap map = rm.getMap()
	//map.moveEnvironment("trap" + rm.getReactor().getName()[5], 1)
}
