package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New2519000() script.Script {
	return generic.NewReactor(2519000, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		//int denyWidth = 320, denyHeight = 150
		//Point denyPos = rm.getReactor().position()
		//Rectangle denyArea = new Rectangle((denyPos.getX() - denyWidth / 2).intValue(), (denyPos.getY() - denyHeight / 2).intValue(), denyWidth, denyHeight)
		//
		//rm.getReactor().getMap().setAllowSpawnPointInBox(false, denyArea)
		//
		//MapleMap map = rm.getReactor().getMap()
		//if (map.getReactorByName("sMob2").getState() >= 1 && map.getReactorByName("sMob3").getState() >= 1 && map.getReactorByName("sMob4").getState() >= 1 && map.countMonsters() == 0) {
		//	rm.getEventInstance().showClearEffect(map.getId())
		//}
	}))
}
