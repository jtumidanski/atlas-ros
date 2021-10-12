package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"math/rand"
	"strconv"
)

func NewTransparent(id uint32) script.Script {
	return generic.NewReactor(id, generic.SetAct(TransparentAct))
}

func TransparentAct(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
	if !event.ParticipatingInEvent(l)(c.CharacterId) {
		return
	}

	e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
	if err != nil {
		return
	}

	if event.GetProperty(l)(e.Id(), "statusStg2") == -1 {
		rnd := int64(math.Max(math.Floor(rand.Float64()*14), 4))
		event.SetStringProperty(l)(e.Id(), "statusStg2", strconv.FormatInt(rnd, 10))
		event.SetStringProperty(l)(e.Id(), "statusStg2_c", "0")
	}
	limit := event.GetProperty(l)(e.Id(), "statusStg2")
	count := event.GetProperty(l)(e.Id(), "statusStg2_c")
	if count >= limit {
		generic.Drop(false, 0, 0, 0, 0)(l, span, db, c)
		event.GiveParticipantsExperience(l)(e.Id(), 3500)
		event.SetStringProperty(l)(e.Id(), "statusStg2", "1")
		generic.ShowClearEffectWithGate(true)(l, db, c)
		return
	}

	count++
	event.SetProperty(l)(e.Id(), "statusStg2_c", count)
	nh := uint32((11 * count) % 14)
	r, err := reactor.GetByClassificationInMap(c.WorldId, c.ChannelId, c.MapId, 2001002+nh)
	if err != nil {
		return
	}
	generic.SpawnMonsterAt(9300040, r.X(), r.Y())
}
