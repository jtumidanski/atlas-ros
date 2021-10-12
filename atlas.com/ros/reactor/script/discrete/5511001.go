package discrete

import (
	_map "atlas-ros/map"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func New5511001() script.Script {
	return generic.NewReactor(5511001, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if _map.MonsterCount(l)(c.WorldId, c.ChannelId, c.MapId, 9420547) == 0 {
			go func() {
				time.Sleep(3200 * time.Millisecond)
				generic.SpawnMonsterAt(9420547, -238, 636)(l, span, db, c)
				generic.ChangeMusic("Bgm09/TimeAttack")(l, span, db, c)
				generic.MapBlueMessage("SCARLION_SUMMONED")(l, span, db, c)
			}()
		}
	}))
}
