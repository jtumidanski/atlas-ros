package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
)

func New2406000() script.Script {
	return generic.NewReactor(2406000,
		generic.SetActs(generic.SpawnNPC(2081008), generic.StartQuest(100203), generic.MapBlueMessage("BABY_DRAGON_SUMMONED")),
		generic.SetTouches(generic.HitReactor(), generic.GainItem(4001094, -1)))
}
