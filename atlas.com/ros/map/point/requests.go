package point

import (
	"atlas-ros/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	mapInformationServicePrefix string = "/ms/mis/"
	mapInformationService              = requests.BaseRequest + mapInformationServicePrefix
	mapResource                        = mapInformationService + "maps/%d"
	dropPosition                       = mapResource + "/dropPosition"
)

func CalculateDropPosition(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, initialX int16, initialY int16, fallbackX int16, fallbackY int16) (requests.DataBody[MapPointAttributes], error) {
	return func(mapId uint32, initialX int16, initialY int16, fallbackX int16, fallbackY int16) (requests.DataBody[MapPointAttributes], error) {
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
		ar, _, err := requests.MakePostRequest[MapPointAttributes](fmt.Sprintf(dropPosition, mapId), input)(l, span)
		if err != nil {
			return requests.DataBody[MapPointAttributes]{}, err
		}
		return ar.Data(), nil
	}
}
