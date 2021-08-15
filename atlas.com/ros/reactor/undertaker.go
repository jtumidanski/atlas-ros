package reactor

import (
	"atlas-ros/configuration"
	"github.com/sirupsen/logrus"
	"time"
)

type Undertaker struct {
	l        logrus.FieldLogger
	interval time.Duration
	timeout  time.Duration
}

func NewUndertaker(l logrus.FieldLogger, interval time.Duration) *Undertaker {
	var to int64
	c, err := configuration.GetConfiguration()
	if err != nil {
		to = 60000
	} else {
		to = c.UndertakerDuration
	}

	timeout := time.Duration(to) * time.Millisecond
	l.Infof("Initializing timeout task to run every %dms, timeout session older than %dms", interval.Milliseconds(), timeout.Milliseconds())
	return &Undertaker{l, interval, timeout}
}

func (t *Undertaker) Run() {
	reactors := GetRegistry().GetAllDead()
	cur := time.Now()

	for _, r := range reactors {
		if cur.Sub(r.UpdateTime()) > t.timeout {
			t.l.Infof("Reactor %d was culled by the undertaker task.", r.Id())
			GetRegistry().Remove(r.Id())
		}
	}
}

func (t *Undertaker) SleepTime() time.Duration {
	return t.interval
}
