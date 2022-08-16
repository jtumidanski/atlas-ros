package drop

type Model struct {
	itemId  uint32
	questId int32
	chance  uint32
}

func (m Model) Chance() uint32 {
	return m.chance
}

func (m Model) ItemId() uint32 {
	return m.itemId
}

func (m Model) QuestId() int32 {
	return m.questId
}

func NewMesoModel(chance uint32) Model {
	return Model{
		itemId:  0,
		questId: -1,
		chance:  chance,
	}
}
