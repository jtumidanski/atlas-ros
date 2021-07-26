package generic

import (
	"atlas-ros/drop"
	"atlas-ros/reactor/script"
	"github.com/sirupsen/logrus"
)

type config struct {
	act     script.ActFunc
	hit     script.HitFunc
	touch   script.TouchFunc
	release script.ReleaseFunc
}

type Configurator func(c *config)

type reactor struct {
	id uint32
	c  config
}

func (r reactor) ReactorId() uint32 {
	return r.id
}

func (r reactor) Act(l logrus.FieldLogger, c script.Context) {
	r.c.act(l, c)
}

func (r reactor) Hit(l logrus.FieldLogger, c script.Context) {
	r.c.hit(l, c)
}

func (r reactor) Touch(l logrus.FieldLogger, c script.Context) {
	r.c.touch(l, c)
}

func (r reactor) Release(l logrus.FieldLogger, c script.Context) {
	r.c.release(l, c)
}

func NewNoOpReactor(id uint32) script.Script {
	return reactor{
		id: id,
		c: config{
			act:     NoOp,
			hit:     NoOp,
			touch:   NoOp,
			release: NoOp,
		},
	}
}

func NewReactor(id uint32, configurators ...Configurator) script.Script {
	c := &config{
		act:     SimpleDrop(false, 0, 0, 0, 0),
		hit:     NoOp,
		touch:   NoOp,
		release: NoOp,
	}
	for _, a := range configurators {
		a(c)
	}
	return reactor{
		id: id,
		c:  *c,
	}
}

func SimpleDropReactor(id uint32, meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) script.Script {
	return NewReactor(id, SetAct(SimpleDrop(meso, mesoChance, minMeso, maxMeso, minItems)))
}

func SimpleWarpByNameReactor(id uint32, mapId uint32, portalName string) script.Script {
	return NewReactor(id, SetAct(SimpleWarpByName(mapId, portalName)))
}

func SimpleSpawnMonsterReactor(id uint32, monsterId uint32) script.Script {
	return NewReactor(id, SetAct(SimpleSpawnMonster(monsterId)))
}

func SimpleSprayReactor(id uint32, meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) script.Script {
	return NewReactor(id, SetAct(SimpleSpray(meso, mesoChance, minMeso, maxMeso, minItems)))
}

func SetAct(actFunc script.ActFunc) Configurator {
	return func(c *config) {
		c.act = actFunc
	}
}

func SetHit(hitFunc script.HitFunc) Configurator {
	return func(c *config) {
		c.hit = hitFunc
	}
}

func SimpleDrop(meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) script.ActFunc {
	return func(l logrus.FieldLogger, c script.Context) {
		drop.Produce(l)(c.WorldId, c.ChannelId, c.MapId, c.ReactorId, c.CharacterId, meso, mesoChance, minMeso, maxMeso, minItems)
	}
}

func SimpleSpray(meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) script.ActFunc {
	return func(l logrus.FieldLogger, c script.Context) {

	}
}

func SimpleWarpById(mapId uint32, portalId uint32) script.ActFunc {
	return func(l logrus.FieldLogger, c script.Context) {
		//TODO
	}
}

func SimpleWarpByName(mapId uint32, portalName string) script.ActFunc {
	return func(l logrus.FieldLogger, c script.Context) {
		//TODO
	}
}

func SimpleSpawnMonster(monsterId uint32) script.ActFunc {
	return func(l logrus.FieldLogger, c script.Context) {

	}
}

func NoOp(_ logrus.FieldLogger, _ script.Context) {
}
