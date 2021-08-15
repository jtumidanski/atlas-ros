package reactor

import (
	"atlas-ros/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func InitResource(router *mux.Router, l logrus.FieldLogger, db *gorm.DB) {
	r := router.PathPrefix("/reactors").Subrouter()
	r.HandleFunc("/{id}", ParseId(l, db, HandleGetReactor)).Methods(http.MethodGet)
	r.HandleFunc("/{id}/reset", ParseId(l, db, HandleResetReactor)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", ParseId(l, db, HandleUpdateReactor)).Methods(http.MethodPatch)
	r.HandleFunc("/{id}", ParseId(l, db, HandleDestroyReactor)).Methods(http.MethodDelete)
	r.HandleFunc("/{id}/hits/relationships/characters", ParseId(l, db, HandleHitReactor)).Methods(http.MethodPost)
}

type IdHandler func(l logrus.FieldLogger, db *gorm.DB, reactorId uint32) http.HandlerFunc

func ParseId(l logrus.FieldLogger, db *gorm.DB, next IdHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reactorId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			l.WithError(err).Errorf("Unable to properly parse reactorId from path.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next(l, db, uint32(reactorId))(w, r)
	}
}

func HandleGetReactor(l logrus.FieldLogger, _ *gorm.DB, reactorId uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		r, err := GetById(reactorId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		result := &DataContainer{Data: MakeReactorBody(*r)}

		w.WriteHeader(http.StatusOK)
		err = json.ToJSON(result, w)
		if err != nil {
			l.WithError(err).Errorf("Encoding response")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func HandleResetReactor(_ logrus.FieldLogger, _ *gorm.DB, _ uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func HandleUpdateReactor(_ logrus.FieldLogger, _ *gorm.DB, _ uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func HandleDestroyReactor(_ logrus.FieldLogger, _ *gorm.DB, _ uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func HandleHitReactor(_ logrus.FieldLogger, _ *gorm.DB, _ uint32) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
