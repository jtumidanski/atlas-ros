package discrete

import (
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"math/rand"
)

func NewStoneShieldSeiramsShield() script.Script {
	return generic.NewReactor(reactor.StoneShield, generic.SetAct(StoneShieldSeiramsShieldAct))
}

func StoneShieldSeiramsShieldAct(l logrus.FieldLogger, c script.Context) {
	if rand.Float64() > 0.7 {
		generic.SimpleDrop(false, 0, 0, 0, 0)(l, c)
	} else {
		generic.SimpleWarpById(_map.DungeonAnotherEntrance, 0)(l, c)
	}
}
