package _map

import (
	"atlas-ros/json"
	"atlas-ros/kafka/producers"
	"atlas-ros/reactor"
	"atlas-ros/rest/resource"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func InitResource(router *mux.Router, l logrus.FieldLogger, db *gorm.DB) {
	w := router.PathPrefix("/worlds").Subrouter()
	w.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/reactors", ParseMap(l, db, HandleGetReactors)).Methods(http.MethodGet)
	w.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/reactors", ParseMap(l, db, HandleCreateReactor)).Methods(http.MethodPost)
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

func HandleCreateReactor(l logrus.FieldLogger, _ *gorm.DB, worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		li := &reactor.InputDataContainer{}
		err := json.FromJSON(li, r.Body)
		if err != nil {
			l.WithError(err).Errorf("Deserializing input.")
			w.WriteHeader(http.StatusBadRequest)
			err = json.ToJSON(&resource.GenericError{Message: err.Error()}, w)
			if err != nil {
				l.WithError(err).Fatalf("Writing error message.")
			}
			return
		}
		attr := li.Data.Attributes
		producers.CreateReactor(l)(worldId, channelId, mapId, attr.Id, attr.Name, attr.State, attr.X, attr.Y, attr.Delay, attr.FacingDirection)
		w.WriteHeader(http.StatusAccepted)
	}
}

func HandleGetReactors(l logrus.FieldLogger, _ *gorm.DB, worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		reactors := reactor.GetInMap(l)(worldId, channelId, mapId)

		result := &reactor.DataListContainer{Data: make([]reactor.DataBody, 0)}
		for _, r := range reactors {
			body := reactor.MakeReactorBody(r)
			result.Data = append(result.Data, body)
		}

		err := json.ToJSON(result, w)
		if err != nil {
			l.WithError(err).Errorf("Encoding response")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
