package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewAltar() script.Script {
	return generic.NewReactor(reactor.Altar, generic.SetAct(AltarAct))
}

func AltarAct(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	if event.ParticipatingInEvent(l)(c.CharacterId) {
		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}
		event.SetStringProperty(l)(e.Id(), "summoned", "true")
		event.SetStringProperty(l)(e.Id(), "canEnter", "false")
	}

	generic.ChangeMusic("Bgm06/FinalFight")(l, db, c)
	generic.SpawnFakeMonster(8800000)(l, db, c)
	for i := uint32(8800003); i < 8800011; i++ {
		generic.SpawnMonster(i)(l, db, c)
	}
	generic.CreateMapMonitor(280030000, "ps00")(l, db, c)
	generic.MapPinkMessage("ZAKUM_SUMMONED")(l, db, c)
}
