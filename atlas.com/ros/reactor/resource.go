package reactor

import (
	"atlas-ros/json"
	"atlas-ros/rest"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	getReactor     = "get_reactor"
	resetReactor   = "reset_reactor"
	updateReactor  = "update_reactor"
	destroyReactor = "destroy_reactor"
	hitReactor     = "hit_reactor"
)

func InitResource(router *mux.Router, l logrus.FieldLogger, db *gorm.DB) {
	r := router.PathPrefix("/reactors").Subrouter()
	r.HandleFunc("/{id}", registerGetReactor(l, db)).Methods(http.MethodGet)
	r.HandleFunc("/{id}/reset", registerResetReactor(l, db)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", registerUpdateReactor(l, db)).Methods(http.MethodPatch)
	r.HandleFunc("/{id}", registerDestroyReactor(l, db)).Methods(http.MethodDelete)
	r.HandleFunc("/{id}/hits/relationships/characters", registerHitReactor(l, db)).Methods(http.MethodPost)
}

func registerHitReactor(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(hitReactor, func(span opentracing.Span) http.HandlerFunc {
		return ParseId(l, func(reactorId uint32) http.HandlerFunc {
			return handleHitReactor(l, db)(span)(reactorId)
		})
	})
}

func registerDestroyReactor(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(destroyReactor, func(span opentracing.Span) http.HandlerFunc {
		return ParseId(l, func(reactorId uint32) http.HandlerFunc {
			return handleDestroyReactor(l, db)(span)(reactorId)
		})
	})
}

func registerUpdateReactor(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(updateReactor, func(span opentracing.Span) http.HandlerFunc {
		return ParseId(l, func(reactorId uint32) http.HandlerFunc {
			return handleUpdateReactor(l, db)(span)(reactorId)
		})
	})
}

func registerResetReactor(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(resetReactor, func(span opentracing.Span) http.HandlerFunc {
		return ParseId(l, func(reactorId uint32) http.HandlerFunc {
			return handleResetReactor(l, db)(span)(reactorId)
		})
	})
}

func registerGetReactor(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(getReactor, func(span opentracing.Span) http.HandlerFunc {
		return ParseId(l, func(reactorId uint32) http.HandlerFunc {
			return handleGetReactor(l, db)(span)(reactorId)
		})
	})
}

type IdHandler func(reactorId uint32) http.HandlerFunc

func ParseId(l logrus.FieldLogger, next IdHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reactorId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			l.WithError(err).Errorf("Unable to properly parse reactorId from path.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next(uint32(reactorId))(w, r)
	}
}

func handleGetReactor(l logrus.FieldLogger, _ *gorm.DB) func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
		return func(reactorId uint32) http.HandlerFunc {
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
	}
}

func handleResetReactor(_ logrus.FieldLogger, _ *gorm.DB) func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
		return func(reactorId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
			}
		}
	}
}

func handleUpdateReactor(_ logrus.FieldLogger, _ *gorm.DB) func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
		return func(reactorId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
			}
		}
	}
}

func handleDestroyReactor(_ logrus.FieldLogger, _ *gorm.DB) func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
		return func(reactorId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
			}
		}
	}
}

func handleHitReactor(_ logrus.FieldLogger, _ *gorm.DB) func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(reactorId uint32) http.HandlerFunc {
		return func(reactorId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
			}
		}
	}
}
