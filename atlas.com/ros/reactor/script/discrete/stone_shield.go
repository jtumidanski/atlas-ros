package discrete

import (
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
)

func NewStoneShieldSeiramsShield() script.Script {
	return generic.NewReactor(reactor.StoneShield, generic.SetAct(StoneShieldSeiramsShieldAct))
}

func StoneShieldSeiramsShieldAct(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	if rand.Float64() > 0.7 {
		generic.Drop(false, 0, 0, 0, 0)(l, db, c)
	} else {
		generic.WarpById(_map.DungeonAnotherEntrance, 0)(l, db, c)
	}
}
