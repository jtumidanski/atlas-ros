package discrete

import (
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
)

func NewStoneShieldSeiramsShield() script.Script {
	return generic.NewReactor(reactor.StoneShield, generic.SetAct(StoneShieldSeiramsShieldAct))
}

func StoneShieldSeiramsShieldAct(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
	if rand.Float64() > 0.7 {
		generic.Drop(false, 0, 0, 0, 0)(l, span, db, c)
	} else {
		generic.WarpById(_map.DungeonAnotherEntrance, 0)(l, span, db, c)
	}
}
