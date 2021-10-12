package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Touch2408003() script.TouchFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		if event.ParticipatingInEvent(l)(c.CharacterId) {
			e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
			if err != nil {
				return
			}
			event.SetStringProperty(l)(e.Id(), "summoned", "true")
			event.SetStringProperty(l)(e.Id(), "canEnter", "false")
		}
		generic.SpawnFakeMonster(8800000)(l, db, c)
		generic.MapBlueMessage("GIGANTIC_CREATURE")(l, span, db, c)
		generic.CreateMapMonitor(c.MapId, "ps00")(l, span, db, c)
		switch c.MapId {
		case 240060000:
			generic.SpawnMonsterAt(8810000, 960, 0)(l, span, db, c)
			break
		case 240060100:
			//TODO needs correct positions
			generic.SpawnMonsterAt(8810001, 960, 0)(l, span, db, c)
			break
		}
	}
}
