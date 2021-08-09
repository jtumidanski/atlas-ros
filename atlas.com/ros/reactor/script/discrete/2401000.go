package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
)

func Act2401000() []script.ActFunc {
	results := make([]script.ActFunc, 0)
	results = append(results, generic.ChangeMusic("Bgm14/HonTale"))
	results = append(results, generic.SpawnHorntailAt(71, 260))
	results = append(results, generic.RestartEventTimer(60*60000))
	results = append(results, generic.MapPinkMessage("HORN_TAIL_SUMMONED"))
	return results
}
