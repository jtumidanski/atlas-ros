package rest

import (
	_map "atlas-ros/map"
	"atlas-ros/reactor"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
		w.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/reactors", ParseMap(l, db, _map.HandleGetReactors)).Methods(http.MethodGet)
		w.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/reactors", ParseMap(l, db, _map.HandleCreateReactor)).Methods(http.MethodPost)

		return router
	}
}

type MapHandler func(l logrus.FieldLogger, db *gorm.DB, worldId byte, channelId byte, mapId uint32) http.HandlerFunc

func ParseMap(l logrus.FieldLogger, db *gorm.DB, next MapHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		worldId, err := strconv.Atoi(mux.Vars(r)["worldId"])
		if err != nil {
			l.WithError(err).Errorf("Unable to properly parse worldId from path.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		channelId, err := strconv.Atoi(mux.Vars(r)["channelId"])
		if err != nil {
			l.WithError(err).Errorf("Unable to properly parse channelId from path.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		mapId, err := strconv.Atoi(mux.Vars(r)["mapId"])
		if err != nil {
			l.WithError(err).Errorf("Unable to properly parse mapId from path.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next(l, db, byte(worldId), byte(channelId), uint32(mapId))(w, r)
	}
}
