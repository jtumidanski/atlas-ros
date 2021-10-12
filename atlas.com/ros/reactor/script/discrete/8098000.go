package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
)

func New8098000() script.Script {
	return generic.NewReactor(8098000, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		b := uint32(math.Abs(float64(c.MapId) - 809050005))
		if c.MapId != 809050000 && c.MapId != 809050010 && c.MapId != 809050014 {
			generic.SpawnMonsters(9400217-b, 2)(l, span, db, c)
			generic.SpawnMonsters(9400218-b, 3)(l, span, db, c)
		} else {
			generic.SpawnMonsters(9400209-b, 6)(l, span, db, c)
			generic.SpawnMonsters(9400210-b, 9)(l, span, db, c)
		}
		generic.MapPinkMessage("SOME_MONSTERS_SUMMONED")(l, span, db, c)
	}))
}
