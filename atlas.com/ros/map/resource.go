package _map

import (
	"atlas-ros/json"
	"atlas-ros/reactor"
	"atlas-ros/rest"
	"atlas-ros/rest/resource"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	getReactors   = "get_reactors"
	createReactor = "create_reactor"
)

func InitResource(router *mux.Router, l logrus.FieldLogger, _ *gorm.DB) {
	w := router.PathPrefix("/worlds").Subrouter()
	w.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/reactors", registerGetReactors(l)).Methods(http.MethodGet)
	w.HandleFunc("/{worldId}/channels/{channelId}/maps/{mapId}/reactors", registerCreateReactors(l)).Methods(http.MethodPost)
	//TODO AT-1 implement reactors?name=
	//TODO AT-2 implement reactors?classification=
}

func registerCreateReactors(l logrus.FieldLogger) http.HandlerFunc {
	return rest.RetrieveSpan(createReactor, func(span opentracing.Span) http.HandlerFunc {
		return ParseMap(l, func(worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
			return handleCreateReactor(l)(span)(worldId, channelId, mapId)
		})
	})
}

type MapHandler func(worldId byte, channelId byte, mapId uint32) http.HandlerFunc

func ParseMap(l logrus.FieldLogger, next MapHandler) http.HandlerFunc {
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
		next(byte(worldId), byte(channelId), uint32(mapId))(w, r)
	}
}

func handleCreateReactor(l logrus.FieldLogger) func(span opentracing.Span) func(worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
		return func(worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
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
				reactor.DeferCreate(l, span)(worldId, channelId, mapId, attr.Classification, attr.Name, attr.State, attr.X, attr.Y, attr.Delay, attr.FacingDirection)
				w.WriteHeader(http.StatusAccepted)
			}
		}
	}
}

func registerGetReactors(l logrus.FieldLogger) http.HandlerFunc {
	return rest.RetrieveSpan(getReactors, func(span opentracing.Span) http.HandlerFunc {
		return ParseMap(l, func(worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
			return handleGetReactors(l)(span)(worldId, channelId, mapId)
		})
	})
}

func handleGetReactors(l logrus.FieldLogger) func(span opentracing.Span) func(worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
		return func(worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, _ *http.Request) {
				reactors := reactor.GetInMap(l)(worldId, channelId, mapId)

				result := &reactor.DataListContainer{Data: make([]reactor.DataBody, 0)}
				for _, r := range reactors {
					body := reactor.MakeReactorBody(r)
					result.Data = append(result.Data, body)
				}

				w.WriteHeader(http.StatusOK)
				err := json.ToJSON(result, w)
				if err != nil {
					l.WithError(err).Errorf("Encoding response")
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
	}
}
