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

func New6109013() script.Script {
	return generic.NewReactor(6109013, generic.SetAct(generic.NoOp), generic.SetTouch(Touch6109013()), generic.SetRelease(Release6109013()))
}

func Release6109013() script.ReleaseFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		fid := "glpq_s"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 5 {
			Action6109013(l, c, e.Id())
		}
		event.SetProperty(l)(e.Id(), fid, cur+1)
	}
}

func Action6109013(l logrus.FieldLogger, c script.Context, eventId uint32) {
	generic.MapBlueMessage("STIRGES_DISAPPEARED")
	_map.KillAllMonsters(l)(c.WorldId, c.ChannelId, c.MapId)
	event.SetProperty(l)(eventId, "glpq_s", 777)
}

func Touch6109013() script.TouchFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		fid := "glpq_s"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 5 {
			Action6109013(l, c, e.Id())
		}
		event.SetProperty(l)(e.Id(), fid, cur-1)
	}
}
