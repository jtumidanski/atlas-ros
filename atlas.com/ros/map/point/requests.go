package point

import (
	"atlas-ros/rest/requests"
	"fmt"
)

const (
	mapInformationServicePrefix string = "/ms/mis/"
	mapInformationService              = requests.BaseRequest + mapInformationServicePrefix
	mapResource                        = mapInformationService + "maps/%d"
	dropPosition                       = mapResource + "/dropPosition"
)

func CalculateDropPosition(mapId uint32, initialX int16, initialY int16, fallbackX int16, fallbackY int16) (*MapPointDataContainer, error) {
	input := &DropPositionInputDataContainer{Data: DropPositionData{
		Id:   "0",
		Type: "com.atlas.mis.attribute.DropPositionInputAttributes",
		Attributes: DropPositionAttributes{
			InitialX:  initialX,
			InitialY:  initialY,
			FallbackX: fallbackX,
			FallbackY: fallbackY,
		},
	}}
	resp, err := requests.Post(fmt.Sprintf(dropPosition, mapId), input)
	if err != nil {
		return nil, err
	}

	ar := &MapPointDataContainer{}
	err = requests.ProcessResponse(resp, ar)
	if err != nil {
		return nil, err
	}

	return ar, nil
}
