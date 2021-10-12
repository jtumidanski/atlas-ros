package discrete

import (
	"atlas-ros/event"
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New9108001() script.Script {
	return generic.NewReactor(9108001, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		r, err := reactor.GetByNameInMap(c.WorldId, c.ChannelId, c.MapId, "fullmoon")
		if err != nil {
			return
		}
		stage := event.GetProperty(l)(e.Id(), "stage") + 1
		event.SetProperty(l)(e.Id(), "stage", stage)
		reactor.TryForceHitReactor(l, span)(r.Id(), r.State() + 1)
		if stage == 6 {
			generic.MapBlueMessage("PROTECT_THE_MOON_BUNNY")(l, span, db, c)
			_map.SetSummonState(l)(c.WorldId, c.ChannelId, c.MapId, true)
			generic.SpawnMonsterAt(9300061, -183, -433)
		}
	}))
}
