package discrete

import (
	_map "atlas-ros/map"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func New5511000() script.Script {
	return generic.NewReactor(5511000, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		if _map.MonsterCount(l)(c.WorldId, c.ChannelId, c.MapId, 9420542) == 0 {
			go func() {
				time.Sleep(3200 * time.Millisecond)
				generic.SpawnMonsterAt(9420542, -527, 637)(l, db, c)
				generic.ChangeMusic("Bgm09/TimeAttack")(l, db, c)
				generic.MapBlueMessage("TARGA_SUMMONED")(l, db, c)
			}()
		}
	}))
}
