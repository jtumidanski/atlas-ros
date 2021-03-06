package reactor

import (
	"strconv"
	"time"
)

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
	WorldId         byte      `json:"world_id"`
	ChannelId       byte      `json:"channel_id"`
	MapId           uint32    `json:"map_id"`
	Classification  uint32    `json:"classification"`
	Name            string    `json:"name"`
	Type            int32     `json:"type"`
	State           int8      `json:"state"`
	EventState      byte      `json:"event_state"`
	X               int16     `json:"x"`
	Y               int16     `json:"y"`
	Delay           uint32    `json:"delay"`
	FacingDirection byte      `json:"facing_direction"`
	Alive           bool      `json:"alive"`
	UpdateTime      time.Time `json:"update_time"`
}

func MakeReactorBody(r Model) DataBody {
	return DataBody{
		Id:   strconv.Itoa(int(r.Id())),
		Type: "reactor",
		Attributes: Attributes{
			WorldId:         r.WorldId(),
			ChannelId:       r.ChannelId(),
			MapId:           r.MapId(),
			Classification:  r.Classification(),
			Name:            r.Name(),
			Type:            r.Type(),
			State:           r.State(),
			EventState:      r.EventState(),
			X:               r.X(),
			Y:               r.Y(),
			Delay:           r.Delay(),
			FacingDirection: r.FacingDirection(),
			Alive:           r.Alive(),
			UpdateTime:      r.UpdateTime(),
		},
	}
}
