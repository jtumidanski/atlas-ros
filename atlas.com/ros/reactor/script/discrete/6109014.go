package discrete

import (
	"atlas-ros/event"
	_map "atlas-ros/map"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New6109014() script.Script {
	return generic.NewReactor(6109014, generic.SetAct(generic.NoOp), generic.SetTouch(Touch6109014()), generic.SetRelease(Release6109014()))
}

func Touch6109014() script.TouchFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		fid := "glpq_f0"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 0 {
			Action6109014(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur+1)
	}
}

func Release6109014() script.ReleaseFunc {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		fid := "glpq_f0"
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		cur := event.GetProperty(l)(e.Id(), fid)
		if cur == 1 {
			Action6109014(l, c)
		}
		event.SetProperty(l)(e.Id(), fid, cur-1)
	}
}

func Action6109014(l logrus.FieldLogger, c script.Context) {
	flames := []string{"a1", "a2", "b1", "b2", "c1", "c2"}
	for i := 0; i < len(flames); i++ {
		_map.ToggleEnvironment(l)(c.WorldId, c.ChannelId, c.MapId, flames[i])
	}
}
