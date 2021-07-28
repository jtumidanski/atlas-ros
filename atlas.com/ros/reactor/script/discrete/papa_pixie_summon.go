package discrete

import (
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
	//rm.getMap().killAllMonsters()
	//rm.getMap().allowSummonState(false)
	//rm.spawnMonster(9300039, 260, 490)
	//MessageBroadcaster.getInstance().sendMapServerNotice(rm.getPlayer().getMap(), ServerNoticeType.PINK_TEXT, I18nMessage.from("2001016_AS_THE_AIR"))
}
