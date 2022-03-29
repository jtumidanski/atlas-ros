package main

import (
	"atlas-ros/database"
	"atlas-ros/kafka"
	"atlas-ros/logger"
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"atlas-ros/reactor/drop"
	"atlas-ros/reactor/script/initializer"
	"atlas-ros/reactor/script/registry"
	"atlas-ros/rest"
	"atlas-ros/tracing"
	"atlas-ros/wz"
	"context"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const serviceName = "atlas-ros"
const consumerGroupId = "Reactor Orchestration Service"

func main() {
	l := logger.CreateLogger(serviceName)
	l.Infoln("Starting main service.")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	tc, err := tracing.InitTracer(l)(serviceName)
	if err != nil {
		l.WithError(err).Fatal("Unable to initialize tracer.")
	}
	defer func(tc io.Closer) {
		err := tc.Close()
		if err != nil {
			l.WithError(err).Errorf("Unable to close tracer.")
		}
	}(tc)

	wzDir := os.Getenv("WZ_DIR")
	wz.GetFileCache().Init(wzDir)

	db := database.Connect(l, database.SetMigrations(drop.Migration))

	drop.Initialize(l, db)

	registry.GetRegistry().AddScripts(initializer.CreateScripts)

	kafka.CreateConsumers(l, ctx, wg,
		reactor.CreateConsumer(db)(consumerGroupId),
		reactor.HitConsumer(db)(consumerGroupId),
		reactor.TouchConsumer(db)(consumerGroupId),
		reactor.ReleaseConsumer(db)(consumerGroupId))

	rest.CreateService(l, db, ctx, wg, "/ms/ros", reactor.InitResource, _map.InitResource)

	//go tasks.Register(reactor.NewUndertaker(l, time.Millisecond*time.Duration(config.UndertakerTaskInterval)))

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
