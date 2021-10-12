package discrete

import (
	"atlas-ros/event"
	_map "atlas-ros/map"
	"atlas-ros/reactor/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Hit2618001() script.HitFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		is := event.GetProperty(l)(e.Id(), "isAlcadno")
		name := "jnr32_out"
		mapId := uint32(926110202)
		if is == 0 {
			name = "rnf32_out"
			mapId = 926100202
		}
		_map.HitReactor(l)(c.WorldId, c.ChannelId, mapId, name)
	}
}
