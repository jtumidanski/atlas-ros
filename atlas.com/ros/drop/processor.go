package drop

import (
	"atlas-ros/character"
	"atlas-ros/item"
	"atlas-ros/map/point"
	"atlas-ros/reactor"
	drop2 "atlas-ros/reactor/drop"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func Produce(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB) func(worldId byte, channelId byte, mapId uint32, reactorId uint32, characterId uint32, meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) {
	return func(worldId byte, channelId byte, mapId uint32, reactorId uint32, characterId uint32, meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) {
		r, err := reactor.GetById(reactorId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve reactor to drop from.")
			return
		}

		dr := character.GetDropRate(l)(characterId)
		di := drop2.GetByClassification(l, db)(r.Classification())
		items := assembleReactorDropEntries(l)(characterId, generateDropList(l)(di, dr, meso, mesoChance, minItems))

		dx := r.X()
		dy := r.Y()
		if len(items)%2 == 0 {
			dx -= 12
		}
		dx -= int16(12 * len(items))

		//TODO AT-8 get meso rate
		mesoRate := float64(1)

		for _, i := range items {
			time.Sleep(200 * time.Millisecond)
			if i.ItemId() == 0 {
				mesoAmount := uint32(((rand.Float64() * float64(maxMeso-minMeso)) + float64(minMeso)) * mesoRate)
				spawnMeso(l, span)(worldId, channelId, mapId, reactorId, dx, dy, r.X(), r.Y(), characterId, 2, mesoAmount)
			} else {
				spawnItem(l, span)(worldId, channelId, mapId, i.ItemId(), reactorId, dx, dy, r.X(), r.Y(), characterId, 0)
			}
			dx += 25
		}
	}
}

func spawnItem(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, itemId uint32, uniqueId uint32, x int16, y int16, dropperX int16, dropperY int16, killerId uint32, dropType byte) {
	return func(worldId byte, channelId byte, mapId uint32, itemId uint32, uniqueId uint32, x int16, y int16, dropperX int16, dropperY int16, killerId uint32, dropType byte) {
		quantity := uint32(1)
		spawnDrop(l, span)(worldId, channelId, mapId, itemId, quantity, 0, x, y, dropperX, dropperY, uniqueId, killerId, false, dropType)
	}
}

func spawnMeso(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, x int16, y int16, dropperX int16, dropperY int16, killerId uint32, dropType byte, amount uint32) {
	return func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, x int16, y int16, dropperX int16, dropperY int16, killerId uint32, dropType byte, amount uint32) {
		spawnDrop(l, span)(worldId, channelId, mapId, 0, 0, amount, x, y, dropperX, dropperY, uniqueId, killerId, false, dropType)
	}
}

func spawnDrop(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, itemId uint32, quantity uint32, mesos uint32, posX int16, posY int16, dropperX int16, dropperY int16, uniqueId uint32, killerId uint32, playerDrop bool, dropType byte) {
	return func(worldId byte, channelId byte, mapId uint32, itemId uint32, quantity uint32, mesos uint32, posX int16, posY int16, dropperX int16, dropperY int16, uniqueId uint32, killerId uint32, playerDrop bool, dropType byte) {
		tempX, tempY := calculateDropPosition(l, span)(mapId, posX, posY, dropperX, dropperY)
		tempX, tempY = calculateDropPosition(l, span)(mapId, tempX, tempY, tempX, tempY)
		Spawn(l, span)(worldId, channelId, mapId, itemId, quantity, mesos, dropType, tempX, tempY, killerId, 0, uniqueId, dropperX, dropperY, playerDrop, byte(1))
	}
}

func calculateDropPosition(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, initialX int16, initialY int16, fallbackX int16, fallbackY int16) (int16, int16) {
	return func(mapId uint32, initialX int16, initialY int16, fallbackX int16, fallbackY int16) (int16, int16) {
		resp, err := point.CalculateDropPosition(l, span)(mapId, initialX, initialY, fallbackX, fallbackY)
		if err != nil {
			return fallbackX, fallbackY
		} else {
			return resp.Data().Attributes.X, resp.Data().Attributes.Y
		}
	}
}

func assembleReactorDropEntries(l logrus.FieldLogger) func(uint32, []*drop2.Model) []*drop2.Model {
	return func(characterId uint32, items []*drop2.Model) []*drop2.Model {
		de := make([]*drop2.Model, 0)
		vqe := make([]*drop2.Model, 0)
		oqe := make([]*drop2.Model, 0)
		for _, d := range items {
			if !item.QuestItem(l)(d.ItemId()) {
				de = append(de, d)
				continue
			}
			if character.NeedsQuestItem(l)(characterId, d.ItemId(), d.QuestId()) {
				vqe = append(vqe, d)
				continue
			}
			oqe = append(oqe, d)
		}
		rand.Shuffle(len(de), func(i, j int) { de[i], de[j] = de[j], de[i] })
		rand.Shuffle(len(vqe), func(i, j int) { vqe[i], vqe[j] = vqe[j], vqe[i] })
		rand.Shuffle(len(oqe), func(i, j int) { oqe[i], oqe[j] = oqe[j], oqe[i] })

		ir := make([]*drop2.Model, 0)
		ir = append(ir, de...)
		ir = append(ir, vqe...)
		ir = append(ir, oqe...)

		left := make([]*drop2.Model, 0)
		right := make([]*drop2.Model, 0)
		for i := range ir {
			if i%2 == 0 {
				left = append(left, ir[i])
			} else {
				right = append(right, ir[i])
			}
		}

		// reverse the ordering of left
		for i, j := 0, len(left)-1; i < j; i, j = i+1, j-1 {
			left[i], left[j] = left[j], left[i]
		}

		results := make([]*drop2.Model, 0)
		results = append(results, left...)
		results = append(results, right...)
		return results
	}
}

func generateDropList(_ logrus.FieldLogger) func([]*drop2.Model, float64, bool, uint32, uint32) []*drop2.Model {
	return func(drops []*drop2.Model, dropRate float64, meso bool, mesoChance uint32, minItems uint32) []*drop2.Model {
		results := make([]*drop2.Model, 0)
		if meso && rand.Float64() < 1.0/float64(mesoChance) {
			results = append(results, drop2.NewMesoModel(mesoChance))
		}
		for _, d := range drops {
			if rand.Float64() < (dropRate / float64(d.Chance())) {
				results = append(results, d)
			}
		}
		for uint32(len(results)) < minItems {
			results = append(results, drop2.NewMesoModel(mesoChance))
		}
		return results
	}
}
