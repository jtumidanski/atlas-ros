package discrete

import (
	"atlas-ros/event"
	_map "atlas-ros/map"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New6109021() script.Script {
	return generic.NewReactor(6109021, generic.SetAct(generic.NoOp), generic.SetTouch(Touch6109021()), generic.SetRelease(Release6109021()))
}

func Touch6109021() script.TouchFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		fid := "glpq_f1"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 0 {
			Action6109021(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur+1)
	}
}

func Action6109021(l logrus.FieldLogger, c script.Context) {
	flames := []string{"a3", "a4", "a5", "b3", "b4", "b5", "c3", "c4", "c5"}
	for i := 0; i < len(flames); i++ {
		_map.ToggleEnvironment(l)(c.WorldId, c.ChannelId, c.MapId, flames[i])
	}
}

func Release6109021() script.ReleaseFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		fid := "glpq_f1"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 1 {
			Action6109021(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur-1)
	}
}
