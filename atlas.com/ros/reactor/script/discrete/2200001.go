package discrete

import (
	"atlas-ros/reactor/script"
	"atlas-ros/reactor/script/generic"
	"math/rand"
)

func Warp2200001() script.ActFunc {
	mapId := uint32(922000021)
	if rand.Float64() < 0.5 {
		mapId = 922000020
	}
	return generic.WarpById(mapId, 0)
}
