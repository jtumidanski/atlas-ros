package item

import "github.com/sirupsen/logrus"

func QuestItem(_ logrus.FieldLogger) func(itemId uint32) bool {
	return func(itemId uint32) bool {
		return false
	}
}
