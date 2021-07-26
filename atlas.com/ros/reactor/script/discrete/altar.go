package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
)

func NewAltar() script.Script {
	return generic.NewReactor(reactor.Altar, generic.SetAct(AltarAct))
}

func AltarAct(l logrus.FieldLogger, c script.Context) {
	//if (rm.getPlayer().getEventInstance() != null) {
	//	rm.getPlayer().getEventInstance().setProperty("summoned", "true")
	//	rm.getPlayer().getEventInstance().setProperty("canEnter", "false")
	//}
	//rm.changeMusic("Bgm06/FinalFight")
	//rm.spawnFakeMonster(8800000)
	//for (int i = 8800003; i < 8800011; i++) {
	//	rm.spawnMonster(i)
	//}
	//rm.createMapMonitor(280030000, "ps00")
	//MessageBroadcaster.getInstance().sendMapServerNotice(rm.getPlayer().getMap(), ServerNoticeType.PINK_TEXT, I18nMessage.from("ZAKUM_SUMMONED"))
}
