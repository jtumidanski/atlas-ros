package consumers

import (
	"atlas-ros/kafka/handler"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

func CreateEventConsumers(l *logrus.Logger, db *gorm.DB, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_CREATE_REACTOR_COMMAND", EmptyCreateReactorCommand(), HandleCreateReactorCommand(db))
	cec("TOPIC_HIT_REACTOR_COMMAND", EmptyHitReactorCommand(), HandleHitReactorCommand(db))
	cec("TOPIC_TOUCH_REACTOR_COMMAND", EmptyTouchReactorCommand(), HandleTouchReactorCommand(db))
	cec("TOPIC_RELEASE_REACTOR_COMMAND", EmptyReleaseReactorCommand(), HandleReleaseReactorCommand(db))
}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, topicToken, "Reactor Orchestration Service", emptyEventCreator, processor)
}