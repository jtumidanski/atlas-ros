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

func New6109023() script.Script {
	return generic.NewReactor(6109023, generic.SetAct(generic.NoOp), generic.SetTouch(Touch6109023()), generic.SetRelease(Release6109023()))
}

func Touch6109023() script.TouchFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		fid := "glpq_f3"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 0 {
			Action6109023(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur+1)
	}
}

func Action6109023(l logrus.FieldLogger, c script.Context) {
	flames := []string{"d1", "d2", "e1", "e2", "f1", "f2"}
	for i := 0; i < len(flames); i++ {
		_map.ToggleEnvironment(l)(c.WorldId, c.ChannelId, c.MapId, flames[i])
	}
}

func Release6109023() script.ReleaseFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		fid := "glpq_f3"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 1 {
			Action6109023(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur-1)
	}
}