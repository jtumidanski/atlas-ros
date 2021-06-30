package _map

import (
	"atlas-ros/json"
	"atlas-ros/kafka/producers"
	"atlas-ros/reactor"
	"atlas-ros/rest/resource"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

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

func HandleGetReactors(l logrus.FieldLogger, db *gorm.DB, worldId byte, channelId byte, mapId uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
