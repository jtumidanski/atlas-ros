package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New2110000() script.Script {
	return generic.NewReactor(2110000, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		//MessageBroadcaster.getInstance().sendServerNotice(rm.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("UNKNOWN_FORCE"))
		//rm.warp(280010000, 0)
	}))
}
