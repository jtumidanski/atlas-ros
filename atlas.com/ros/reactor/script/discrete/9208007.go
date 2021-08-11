package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New9208007() script.Script {
	return generic.NewReactor(9208007, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		r, err := reactor.GetByNameInMap(c.WorldId, c.ChannelId, 990000400, "speargate")
		if err != nil {
			return
		}
		reactor.TryForceHitReactor(l)(r.Id(), r.State()+1)
		r, err = reactor.GetById(r.Id())
		if err != nil {
			return
		}
		if r.State() == 4 {
			maps := []uint32{990000400, 990000410, 990000420, 990000430, 990000431, 990000440}
			for i := 0; i < len(maps); i++ {
				generic.ShowClearEffectWithGateAndMap(false, maps[i])(l, db, c)
			}
			generic.GainGuildPoints(20)
		}
	}))
}
