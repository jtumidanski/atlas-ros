package rest

import (
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

func CreateRestService(l *logrus.Logger, db *gorm.DB, ctx context.Context, wg *sync.WaitGroup) {
	go NewServer(l, ctx, wg, ProduceRoutes(db))
}

func ProduceRoutes(db *gorm.DB) func(l logrus.FieldLogger) http.Handler {
	return func(l logrus.FieldLogger) http.Handler {
		router := mux.NewRouter().PathPrefix("/ms/ros").Subrouter().StrictSlash(true)
		router.Use(CommonHeader)

		r := router.PathPrefix("/reactors").Subrouter()
		r.HandleFunc("/{id}", reactor.HandleGetReactor(l, db)).Methods(http.MethodGet)
		r.HandleFunc("/{id}/reset", reactor.HandleResetReactor(l, db)).Methods(http.MethodPost)
		r.HandleFunc("/{id}", reactor.HandleUpdateReactor(l, db)).Methods(http.MethodPatch)
		r.HandleFunc("/{id}", reactor.HandleDestroyReactor(l, db)).Methods(http.MethodDelete)
		r.HandleFunc("/{id}/hits/relationships/characters", reactor.HandleHitReactor(l, db)).Methods(http.MethodPost)

		w := router.PathPrefix("/worlds").Subrouter()
		w.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/reactors", _map.HandleGetReactors(l, db)).Methods(http.MethodGet)
		w.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/reactors", _map.HandleCreateReactor(l, db)).Methods(http.MethodPost)
		w.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/reactors/shuffle", _map.HandleGetReactors(l, db)).Methods(http.MethodPost)


		return router
	}
}
