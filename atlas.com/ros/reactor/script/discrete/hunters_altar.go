package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"math/rand"
)

func NewHuntersAltar() script.Script {
	return generic.NewReactor(reactor.HuntersAltar, generic.SetAct(generic.NoOp), generic.SetHit(HuntersAltarHit))
}

func HuntersAltarHit(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
	//rm.hitMonsterWithReactor(6090001, 4)
	_, _ = reactor.SetEventState(c.ReactorId, byte(math.Floor(rand.Float64()*3)))
}
