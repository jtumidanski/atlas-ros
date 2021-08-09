package discrete

import (
	"atlas-ros/event"
	"atlas-ros/reactor"
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func New9208000() script.Script {
	return generic.NewReactor(9208000, generic.SetAct(func(l logrus.FieldLogger, db *gorm.DB, c script.Context) {
		if !event.ParticipatingInEvent(l)(c.CharacterId) {
			return
		}

		e, err := event.GetByParticipatingCharacter(l)(c.CharacterId)
		if err != nil {
			return
		}

		status := event.GetStringProperty(l)(e.Id(), "stage1status")
		if status == "" || status == "waiting" {
			return
		}

		stage := event.GetProperty(l)(e.Id(), "stage1phase")
		if status == "display" {
			if reactor.IsRecentHitFromAttack(l)(c.ReactorId) {
				return
			}
			prevCombo := event.GetStringProperty(l)(e.Id(), "stage1combo")
			prevCombo += fmt.Sprintf("%03d", int(c.ReactorId%1000))
			event.SetStringProperty(l)(e.Id(), "stage1combo", prevCombo)
			if len(prevCombo) == int(3*(stage+3)) {
				event.SetStringProperty(l)(e.Id(), "stage1status", "active")
				generic.MapPinkMessage("PROCEED_WITH_CAUTION")(l, db, c)
				event.SetStringProperty(l)(e.Id(), "stage1guess", "")
			}
			return
		}

		prevGuess := event.GetStringProperty(l)(e.Id(), "stage1guess")
		if len(prevGuess) != int(3*(stage+3)) {
			prevGuess += fmt.Sprintf("%03d", int(c.ReactorId%1000))
			event.SetStringProperty(l)(e.Id(), "stage1guess", prevGuess)
		}
	}))
}
