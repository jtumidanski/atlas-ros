package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewLumpOfLight() script.Script {
	return generic.NewReactor(reactor.LumpOfLight, generic.SetAct(LumpOfLightAct))
}

func LumpOfLightAct(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	//MessageBroadcaster.getInstance().sendMapServerNotice(rm.getPlayer().getMap(), ServerNoticeType.PINK_TEXT, I18nMessage.from("2006000_AS_THE_LIGHT_FLICKERS"))
	//rm.spawnNpc(2013001)
}
