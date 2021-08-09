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

func New6701000() script.Script {
	return generic.NewReactor(6701000, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		r, err := reactor.GetById(l)(c.ReactorId)
		if err != nil {
			return
		}

		startId := uint32(9400523)
		for i := 0; i < 7; i++ {
			monsterId := startId + uint32(math.Floor(rand.Float64()*3))
			generic.SpawnMonsterAt(monsterId, r.X(), r.Y())
		}
	}))
}
