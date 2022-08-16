package drop

import (
	"atlas-ros/database"
	"atlas-ros/model"
	"gorm.io/gorm"
)

func makeModel(e entity) (Model, error) {
	return Model{
		itemId:  e.ItemId,
		questId: e.QuestId,
		chance:  e.Chance,
	}, nil
}

func getByClassification(classification uint32) database.EntitySliceProvider[entity] {
	return func(db *gorm.DB) model.SliceProvider[entity] {
		return database.SliceQuery[entity](db, &entity{ReactorId: classification})
	}
}
