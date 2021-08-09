package discrete

import (
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewPapaPixieSummon() script.Script {
	return generic.NewReactor(reactor.PapaPixieSummon, generic.SetAct(PapaPixieSummonAct))
}

func PapaPixieSummonAct(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	_map.KillAllMonsters(l)(c.WorldId, c.ChannelId, c.MapId)
	_map.SetSummonState(l)(c.WorldId, c.ChannelId, c.MapId, false)
	generic.SpawnMonsterAt(9300039, 260, 490)(l, db, c)
	generic.MapPinkMessage("2001016_AS_THE_AIR")(l, db, c)
}
