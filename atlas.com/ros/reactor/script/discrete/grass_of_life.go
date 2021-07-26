package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
)

func NewGrassOfLife() script.Script {
	return generic.NewReactor(reactor.GrassOfLife, generic.SetAct(GrassOfLifeAct))
}

func GrassOfLifeAct(l logrus.FieldLogger, c script.Context) {
	//rm.dropItems()
	//
	//EventInstanceManager eim = rm.getEventInstance()
	//eim.setProperty("statusStg7", "1")
}
