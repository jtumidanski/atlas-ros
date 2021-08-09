package generic

import (
	"atlas-ros/character"
	"atlas-ros/drop"
	"atlas-ros/event"
	"atlas-ros/reactor/script"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type config struct {
	act     []script.ActFunc
	hit     []script.HitFunc
	touch   []script.TouchFunc
	release []script.ReleaseFunc
}

type Configurator func(c *config)

type reactor struct {
	id uint32
	c  config
}

func (r reactor) ReactorClassification() uint32 {
	return r.id
}

func (r reactor) Act(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	for _, a := range r.c.act {
		a(l, db, c)
	}
}

func (r reactor) Hit(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	for _, a := range r.c.hit {
		a(l, db, c)
	}
}

func (r reactor) Touch(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	for _, a := range r.c.touch {
		a(l, db, c)
	}
}

func (r reactor) Release(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	for _, a := range r.c.release {
		a(l, db, c)
	}
}

func NewNoOpReactor(id uint32) script.Script {
	return reactor{
		id: id,
		c: config{
			act:     []script.ActFunc{NoOp},
			hit:     []script.HitFunc{NoOp},
			touch:   []script.TouchFunc{NoOp},
			release: []script.ReleaseFunc{NoOp},
		},
	}
}

func NewReactor(id uint32, configurators ...Configurator) script.Script {
	c := &config{
		act:     []script.ActFunc{Drop(false, 0, 0, 0, 0)},
		hit:     []script.HitFunc{NoOp},
		touch:   []script.TouchFunc{NoOp},
		release: []script.ReleaseFunc{NoOp},
	}
	for _, a := range configurators {
		a(c)
	}
	return reactor{
		id: id,
		c:  *c,
	}
}

func NewDropReactor(id uint32, meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) script.Script {
	return NewReactor(id, SetAct(Drop(meso, mesoChance, minMeso, maxMeso, minItems)))
}

func NewWarpByNameReactor(id uint32, mapId uint32, portalName string) script.Script {
	return NewReactor(id, SetAct(WarpByName(mapId, portalName)))
}

func NewSpawnMonsterReactor(id uint32, monsterId uint32) script.Script {
	return NewReactor(id, SetAct(SpawnMonster(monsterId)))
}

func NewSprayReactor(id uint32, meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) script.Script {
	return NewReactor(id, SetAct(Spray(meso, mesoChance, minMeso, maxMeso, minItems)))
}

func SetAct(actFunc script.ActFunc) Configurator {
	return func(c *config) {
		c.act = []script.ActFunc{actFunc}
	}
}

func SetActs(actFuncs ...script.ActFunc) Configurator {
	return func(c *config) {
		c.act = actFuncs
	}
}

func SetHit(hitFunc script.HitFunc) Configurator {
	return func(c *config) {
		c.hit = []script.HitFunc{hitFunc}
	}
}

func SetTouch(touchFunc script.TouchFunc) Configurator {
	return func(c *config) {
		c.touch = []script.TouchFunc{touchFunc}
	}
}

func SetTouches(touchFuncs ...script.TouchFunc) Configurator {
	return func(c *config) {
		c.touch = touchFuncs
	}
}

func SetRelease(releaseFunc script.ReleaseFunc) Configurator {
	return func(c *config) {
		c.release = []script.ReleaseFunc{releaseFunc}
	}
}

func Drop(meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		drop.Produce(l, db)(c.WorldId, c.ChannelId, c.MapId, c.ReactorId, c.CharacterId, meso, mesoChance, minMeso, maxMeso, minItems)
	}
}

func Spray(meso bool, mesoChance uint32, minMeso uint32, maxMeso uint32, minItems uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		//TODO HeavenMS there is no difference between this, and Drop
		drop.Produce(l, db)(c.WorldId, c.ChannelId, c.MapId, c.ReactorId, c.CharacterId, meso, mesoChance, minMeso, maxMeso, minItems)
	}
}

func WarpById(mapId uint32, portalId uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		character.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, portalId)
	}
}

func WarpRandom(mapId uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		character.WarpRandom(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId)
	}
}

func WarpByName(mapId uint32, portalName string) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		character.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, portalName)
	}
}

func SpawnMonster(monsterId uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		//TODO
	}
}

func SpawnFakeMonster(monsterId uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		//TODO
	}
}

func SpawnMonsterAt(monsterId uint32, x int16, y int16) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func SpawnMonsters(monsterId uint32, quantity uint16) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	smf := SpawnMonster(monsterId)
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		for i := uint16(0); i < quantity; i++ {
			smf(l, db, c)
		}
	}
}

func SpawnMonstersAt(monsterId uint32, quantity uint16, x int16, y int16) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	smf := SpawnMonster(monsterId)
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		for i := uint16(0); i < quantity; i++ {
			smf(l, db, c)
		}
	}
}

func PinkMessage(text string) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func MapPinkMessage(text string) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func MapBlueMessage(text string) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func EventBlueMessage(text string) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		event.BlueMessageParticipants(l)(event.ParticipatingCharacterIdProvider(c.CharacterId), text)
	}
}

func ChangeMusic(path string) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func CreateMapMonitor(mapId uint32, portalName string) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func SpawnNPC(npcId uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func GainGuildPoints(amount int16) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func NoOp(_ logrus.FieldLogger, _ *gorm.DB, _ script.Context) {
}

func SetEventProperty(name string, value int32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}
		event.SetProperty(l)(e.Id(), name, value)
	}
}

func SpawnHorntailAt(x int16, y int16) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func RestartEventTimer(amount int32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func StartQuest(questId uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func HitReactor() func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func GainItem(itemId uint32, amount int16) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func ShowClearEffect() func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		event.ShowClearEffect(l)(c.WorldId, c.ChannelId, c.CharacterId, c.MapId)
	}
}

func ShowClearEffectWithGate(hasGate bool, mapId uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func ShowClearEffectWithMapObject(mapObj string, newState uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func ShowClearEffectInMapWithMapObject(mapId uint32, mapObj string, newState uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func GiveEventParticipantsStageReward(stage uint32) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}

func MapCharacterGainExperience(amount int16) func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
	return func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {

	}
}
