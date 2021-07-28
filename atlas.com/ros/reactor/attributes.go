package reactor

import "strconv"

type InputDataContainer struct {
	Data DataBody `json:"data"`
}

type DataListContainer struct {
	Data []DataBody `json:"data"`
}

type DataContainer struct {
	Data DataBody `json:"data"`
}

type DataBody struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Classification  uint32 `json:"classification"`
	Name            string `json:"name"`
	Type            int32  `json:"type"`
	State           int8   `json:"state"`
	EventState      byte   `json:"event_state"`
	X               int16  `json:"x"`
	Y               int16  `json:"y"`
	Delay           uint32 `json:"delay"`
	FacingDirection byte   `json:"facing_direction"`
}

func MakeReactorBody(r Model) DataBody {
	return DataBody{
		Id:   strconv.Itoa(int(r.Id())),
		Type: "reactor",
		Attributes: Attributes{
			Classification:  r.Classification(),
			Name:            r.Name(),
			Type:            r.Type(),
			State:           r.State(),
			EventState:      r.EventState(),
			X:               r.X(),
			Y:               r.Y(),
			Delay:           r.Delay(),
			FacingDirection: r.FacingDirection(),
		},
	}
}
