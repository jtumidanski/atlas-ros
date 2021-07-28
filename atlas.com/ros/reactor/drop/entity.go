package drop

import "gorm.io/gorm"

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(&entity{})
}

type entity struct {
	ID        uint32 `gorm:"primaryKey;autoIncrement;not null"`
	ReactorId uint32 `gorm:"not null"`
	ItemId    uint32 `gorm:"not null"`
	QuestId   int32  `gorm:"not null"`
	Chance    uint32 `gorm:"not null"`
}

func (e entity) TableName() string {
	return "drops"
}
