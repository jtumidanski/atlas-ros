package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New9101000() script.Script {
	return generic.NewReactor(9101000, generic.SetAct(func(l logrus.FieldLogger, span opentracing.Span, db *gorm.DB, c script.Context) {
		//rm.spawnMonster(9300061, 1, 0, 0) // (0, 0) is temp position
		//MapleMap map = rm.getClient().getPlayer().getMap()
		//map.startMapEffect("Protect the Moon Bunny that's pounding the mill, and gather up 10 Moon Bunny's Rice Cakes!", 5120016, 7000)
		//MasterBroadcaster.getInstance().sendToAllInMap(map, new ShowBunny())
		//
		////TODO
		////      rm.getClient().getPlayer().getMap().broadcastMessage(MaplePacketCreator.showHPQMoon());
		////      rm.getClient().getPlayer().getMap().showAllMonsters();
	}))
}
