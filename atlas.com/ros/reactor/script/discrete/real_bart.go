package discrete

import (
	"atlas-ros/character"
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
	if character.QuestStarted(l)(c.CharacterId, 6400) {
		character.SetQuestProgress(l)(c.CharacterId, 6400, 1, 2)
		character.SetQuestProgressString(l)(c.CharacterId, 6400, 6401, "q3")
	}
	generic.PinkMessage("REAL_BART_FOUND")(l, db, c)
}
