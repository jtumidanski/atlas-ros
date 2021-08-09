package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New9980000() script.Script {
	return generic.NewReactor(9980000, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		//rm.dispelAllMonsters((rm.getReactor().getName().substring(1, 2)).toInteger(), (rm.getReactor().getName().substring(0, 1)).toInteger())
	}))
}
