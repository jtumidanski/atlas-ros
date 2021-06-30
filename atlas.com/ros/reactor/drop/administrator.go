package drop

import "gorm.io/gorm"

func create(db *gorm.DB, reactorId uint32, itemId uint32, questId uint32, chance uint32) error {
	a := &entity{
		ReactorId: reactorId,
		ItemId:    itemId,
		QuestId:   questId,
		Chance:    chance,
	}
	return db.Create(a).Error
}

func deleteAll(db *gorm.DB) error {
	return db.Exec("DELETE FROM drops").Error
}
