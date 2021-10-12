package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New2519001() script.Script {
	return generic.NewReactor(2519001, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		//int denyWidth = 320, denyHeight = 150
		//Point denyPos = rm.getReactor().position()
		//Rectangle denyArea = new Rectangle((denyPos.getX() - denyWidth / 2).intValue(), (denyPos.getY() - denyHeight / 2).intValue(), denyWidth, denyHeight)
		//
		//rm.getReactor().getMap().setAllowSpawnPointInBox(false, denyArea)
		//
		//MapleMap map = rm.getReactor().getMap()
		//if (map.getReactorByName("sMob1").getState() >= 1 && map.getReactorByName("sMob3").getState() >= 1 && map.getReactorByName("sMob4").getState() >= 1 && map.countMonsters() == 0) {
		//	rm.getEventInstance().showClearEffect(map.getId())
		//}
	}))
}
