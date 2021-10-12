package consumers

import (
	"atlas-ros/kafka/handler"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

const (
	CreateReactorCommand  = "create_reactor_command"
	HitReactorCommand     = "hit_reactor_command"
	TouchReactorCommand   = "touch_reactor_command"
	ReleaseReactorCommand = "release_reactor_command"
)

func CreateEventConsumers(l *logrus.Logger, db *gorm.DB, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, name string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, name, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_CREATE_REACTOR_COMMAND", CreateReactorCommand, EmptyCreateReactorCommand(), HandleCreateReactorCommand(db))
	cec("TOPIC_HIT_REACTOR_COMMAND", HitReactorCommand, EmptyHitReactorCommand(), HandleHitReactorCommand(db))
	cec("TOPIC_TOUCH_REACTOR_COMMAND", TouchReactorCommand, EmptyTouchReactorCommand(), HandleTouchReactorCommand(db))
	cec("TOPIC_RELEASE_REACTOR_COMMAND", ReleaseReactorCommand, EmptyReleaseReactorCommand(), HandleReleaseReactorCommand(db))
}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, name string, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, name, topicToken, "Reactor Orchestration Service", emptyEventCreator, processor)
}
