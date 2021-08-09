package discrete

import (
	"atlas-ros/event"
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
)

func NewPapaPixiesFlowerPot() script.Script {
	return generic.NewReactor(reactor.PapaPixiesFlowerPot, generic.SetAct(PapaPixiesFlowerPotAct))
}

func PapaPixiesFlowerPotAct(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	if !_map.SummonState(l)(c.WorldId, c.ChannelId, c.MapId) {
		return
	}
	if !event.ParticipatingInEvent(l)(c.CharacterId) {
		return
	}

	e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
	if err != nil {
		return
	}
	count := event.GetProperty(l)(e.Id(), "statusStg7_c")
	if count < 7 {
		monsterId := uint32(9300048)
		if rand.Float64() >= 0.6 {
			monsterId = 9300049
		}
		generic.SpawnMonster(monsterId)(l, db, c)
		event.SetProperty(l)(e.Id(), "statusStg7_c", count+1)
		return
	}
	generic.SpawnMonster(9300049)(l, db, c)
}
