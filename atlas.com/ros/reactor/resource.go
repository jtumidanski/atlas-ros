package reactor

import (
	"atlas-ros/json"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func HandleGetReactor(l logrus.FieldLogger, db *gorm.DB, reactorId uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		r, err := Get(l)(reactorId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		result := &DataContainer{Data: MakeReactorBody(*r)}
		err = json.ToJSON(result, w)
		if err != nil {
			l.WithError(err).Errorf("Encoding response")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func HandleResetReactor(l logrus.FieldLogger, db *gorm.DB, reactorId uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func HandleUpdateReactor(l logrus.FieldLogger, db *gorm.DB, reactorId uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func HandleDestroyReactor(l logrus.FieldLogger, db *gorm.DB, reactorId uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func HandleHitReactor(l logrus.FieldLogger, db *gorm.DB, reactorId uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
