package drop

import (
	"gorm.io/gorm"
)

func listGet(db *gorm.DB, query interface{}) ([]*Model, error) {
	var results []entity
	err := db.Where(query).Find(&results).Error
	if err != nil {
		return nil, err
	}

	var models = make([]*Model, 0)
	for _, e := range results {
		models = append(models, makeModel(&e))
	}
	return models, nil
}

func makeModel(e *entity) *Model {
	return &Model{
		itemId:  e.ItemId,
		questId: e.QuestId,
		chance:  e.Chance,
	}
}

func getByClassification(db *gorm.DB) func(classification uint32) ([]*Model, error) {
	return func(classification uint32) ([]*Model, error) {
		return listGet(db, &entity{ReactorId: classification})
	}
}
