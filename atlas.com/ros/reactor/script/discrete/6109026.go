package discrete

import (
	"atlas-ros/event"
	_map "atlas-ros/map"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New6109026() script.Script {
	return generic.NewReactor(6109026, generic.SetAct(generic.NoOp), generic.SetTouch(Touch6109026()), generic.SetRelease(Release6109026()))
}

func Touch6109026() script.TouchFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		fid := "glpq_f6"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 0 {
			Action6109026(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur+1)
	}
}

func Action6109026(l logrus.FieldLogger, c script.Context) {
	flames := []string{"g3", "g4", "g5", "h3", "h4", "h5", "i3", "i4", "i5"}
	for i := 0; i < len(flames); i++ {
		_map.ToggleEnvironment(l)(c.WorldId, c.ChannelId, c.MapId, flames[i])
	}
}

func Release6109026() script.ReleaseFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		fid := "glpq_f6"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 1 {
			Action6109026(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur-1)
	}
}