package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New2618000() script.Script {
	return generic.NewReactor(2618000, generic.SetAct(generic.NoOp), generic.SetHit(Hit2618000()))
}

func Hit2618000() func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		//if (rm.getReactor().getState() == ((byte) 6)) {
		//EventInstanceManager eim = rm.getEventInstance()
		//
		//int done = eim.getIntProperty("statusStg3") + 1
		//eim.setIntProperty("statusStg3", done)
		//
		//if (done == 3) {
		//eim.showClearEffect()
		//eim.giveEventPlayersStageReward(3)
		//rm.getMap().killAllMonsters()
		//
		//String reactorName = (eim.getIntProperty("isAlcadno") == 0) ? "rnj2_door" : "jnr2_door"
		//rm.getMap().getReactorByName(reactorName).hitReactor(rm.getClient())
		//}
		//}
	}
}
