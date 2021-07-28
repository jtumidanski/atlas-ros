package drop

type JSONObject struct {
	Id    uint32       `json:"id"`
	Items []dropChance `json:"items"`
}

type dropChance struct {
	ItemId  uint32 `json:"item_id"`
	QuestId int32  `json:"quest_id"`
	Chance  uint32 `json:"chance"`
}
