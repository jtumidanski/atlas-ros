package main

import (
	"atlas-ros/configuration"
	"atlas-ros/database"
	"atlas-ros/kafka/consumers"
	"atlas-ros/logger"
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"atlas-ros/reactor/drop"
	"atlas-ros/reactor/script/initializer"
	"atlas-ros/reactor/script/registry"
	"atlas-ros/rest"
	"atlas-ros/tasks"
	"atlas-ros/wz"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	l := logger.CreateLogger()
	l.Infoln("Starting main service.")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	config, err := configuration.GetConfiguration()
	if err != nil {
		l.WithError(err).Fatal("Unable to successfully load configuration.")
	}

	wzDir := os.Getenv("WZ_DIR")
	wz.GetFileCache().Init(wzDir)

	db := database.Connect(l, database.SetMigrations(drop.Migration))

	drop.Initialize(l, db)

	registry.GetRegistry().AddScripts(initializer.CreateScripts)

	consumers.CreateEventConsumers(l, db, ctx, wg)

	rest.CreateService(l, db, ctx, wg, "/ms/ros", reactor.InitResource, _map.InitResource)

	go tasks.Register(reactor.NewUndertaker(l, time.Millisecond*time.Duration(config.UndertakerTaskInterval)))

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infof("Initiating shutdown with signal %s.", sig)
	cancel()
	wg.Wait()
	l.Infoln("Service shutdown.")
}
