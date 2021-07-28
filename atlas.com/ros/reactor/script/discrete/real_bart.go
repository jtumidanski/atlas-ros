package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewRealBart() script.Script {
	return generic.NewReactor(reactor.RealBart, generic.SetAct(RealBartAct))
}

func RealBartAct(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	//if (rm.isQuestStarted(6400)) {
	//	rm.setQuestProgress(6400, 1, 2)
	//	rm.setQuestProgress(6400, 6401, "q3")
	//}
	//MessageBroadcaster.getInstance().sendServerNotice(rm.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("REAL_BART_FOUND"))
}
