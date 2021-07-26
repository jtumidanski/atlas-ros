package discrete

import (
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"github.com/sirupsen/logrus"
)

func NewRecord() script.Script {
	return generic.NewReactor(reactor.Record, generic.SetAct(RecordAct))
}

func RecordAct(l logrus.FieldLogger, c script.Context) {
	// rm.getEventInstance().setProperty("statusStg3", "0")
}
