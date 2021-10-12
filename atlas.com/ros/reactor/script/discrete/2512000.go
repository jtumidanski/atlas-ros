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

func Act2512000() script.ActFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		next := event.GetProperty(l)(e.Id(), "openedBoxes") + 1
		event.SetProperty(l)(e.Id(), "openedBoxes", next)
		generic.Drop(true, 1, 30, 60, 15)(l, span, db, c)
		if _map.MonstersCount(l)(c.WorldId, c.ChannelId, c.MapId) == 0 && GrindModePassed2512000(l, c)(e.Id()) {
			event.ShowClearEffect(l)(c.WorldId, c.ChannelId, c.CharacterId, c.MapId)
		}
	}
}

func GrindModePassed2512000(l logrus.FieldLogger, c script.Context) func(eventId uint32) bool {
	return func(eventId uint32) bool {
		if event.GetProperty(l)(eventId, "grindMode") == 0 {
			return true
		}
		return event.AllReactorsActivatedInMap(l)(c.WorldId, c.ChannelId, c.CharacterId, c.MapId, 2511000, 2517999)
	}
}
