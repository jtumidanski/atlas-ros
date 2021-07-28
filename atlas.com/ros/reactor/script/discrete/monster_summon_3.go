package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewMonsterSummon3() script.Script {
	return generic.NewReactor(reactor.MonsterSummon3, generic.SetAct(MonsterSummon3Act))
}

func MonsterSummon3Act(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	//MessageBroadcaster.getInstance().sendServerNotice(rm.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("MONSTERS_IN_CHEST"))
	//rm.spawnMonster(9300004,3)
}
