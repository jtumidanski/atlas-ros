package discrete

import (
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewLever() script.Script {
	return generic.NewReactor(reactor.Lever, generic.SetAct(generic.NoOp), generic.SetHit(LeverHit))
}

func LeverHit(l logrus.FieldLogger, _ *gorm.DB, c script.Context) {
	r, err := reactor.GetById(c.ReactorId)
	if err != nil {
		return
	}

	_map.MoveEnvironment(l)(c.WorldId, c.ChannelId, c.MapId, "trap" + string([]rune(r.Name())[5]), 1)
}
