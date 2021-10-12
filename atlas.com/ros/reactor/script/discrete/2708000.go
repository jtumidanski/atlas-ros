package discrete

import (
	"atlas-ros/reactor/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Hit2708000() script.HitFunc {
	return func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		//spawnJrBoss(mapObj.getMonsterById(8820019));
		//spawnJrBoss(mapObj.getMonsterById(8820020));
		//spawnJrBoss(mapObj.getMonsterById(8820021));
		//spawnJrBoss(mapObj.getMonsterById(8820022));
		//spawnJrBoss(mapObj.getMonsterById(8820023));
		//mapObj.killMonster(8820000)
	}

	//static def spawnJrBoss(MapleMonster mobObj) {
	//	mobObj.getMap().killMonster(mobObj.id())
	//	int spawnId = mobObj.id() - 17
	//
	//	MapleLifeFactory.getMonster(spawnId).ifPresent({ monster -> mobObj.getMap().spawnMonsterOnGroundBelow(monster, mobObj.position()) })
	//}
}
