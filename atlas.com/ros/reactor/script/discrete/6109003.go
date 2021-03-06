package discrete

import (
	"atlas-ros/event"
	_map "atlas-ros/map"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New6109003() script.Script {
	return generic.NewReactor(6109003, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		if c.MapId == 610030200 {
			generic.EventBlueMessage("THIEF_SIGIL_ACTIVATED")
			next := event.GetProperty(l)(e.Id(), "glpq2") + 1
			event.SetProperty(l)(e.Id(), "glpq2", next)
			if next == 5 {
				generic.EventBlueMessage("ANTELLION_NEXT")
				generic.ShowClearEffectWithMapObject("2pt", 2)
				generic.GiveEventParticipantsStageReward(2)
			}
			return
		}

		if c.MapId == 610030300 {
			generic.EventBlueMessage("THIEF_SIGIL_ACTIVATED_LONG")
			next := event.GetProperty(l)(e.Id(), "glpq3") + 1
			event.SetProperty(l)(e.Id(), "glpq3", next)
			_map.MoveEnvironment(l)(c.WorldId, c.ChannelId, c.MapId, "menhir4", 1)
			if next == 5 && event.GetProperty(l)(e.Id(), "glpq3_p") == 5 {
				generic.EventBlueMessage("ANTELLION_NEXT")
				generic.ShowClearEffectWithMapObject("3pt", 2)
				generic.GiveEventParticipantsStageReward(3)
			}
			return
		}
	}))
}
