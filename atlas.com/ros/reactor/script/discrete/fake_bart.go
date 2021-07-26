package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
)

func NewFakeBart() script.Script {
	return generic.NewReactor(reactor.FakeBart, generic.SetAct(FakeBartAct))
}

func FakeBartAct(l logrus.FieldLogger, c script.Context) {
	//MessageBroadcaster.getInstance().sendServerNotice(rm.getPlayer(), ServerNoticeType.PINK_TEXT, I18nMessage.from("FAILED_TO_FIND_BART"))
	//rm.warp(120000102)
}
