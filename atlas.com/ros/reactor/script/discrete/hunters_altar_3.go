package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"math/rand"
)

func NewHuntersAltar3() script.Script {
	return generic.NewReactor(reactor.HuntersAltar3, generic.SetAct(generic.NoOp), generic.SetHit(HuntersAltar3Hit))
}

func HuntersAltar3Hit(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	//rm.hitMonsterWithReactor(6090001, 4)
	_, _ = reactor.SetEventState(c.ReactorId, byte(math.Floor(rand.Float64()*3)))
}
