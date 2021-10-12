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

func New6109025() script.Script {
	return generic.NewReactor(6109025, generic.SetAct(generic.NoOp), generic.SetTouch(Touch6109025()), generic.SetRelease(Release6109025()))
}

func Touch6109025() script.TouchFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		fid := "glpq_f5"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 0 {
			Action6109025(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur+1)
	}
}

func Action6109025(l logrus.FieldLogger, c script.Context) {
	flames := []string{"g1", "g2", "h1", "h2", "i1", "i2"}
	for i := 0; i < len(flames); i++ {
		_map.ToggleEnvironment(l)(c.WorldId, c.ChannelId, c.MapId, flames[i])
	}
}

func Release6109025() script.ReleaseFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		fid := "glpq_f5"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 1 {
			Action6109025(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur-1)
	}
}