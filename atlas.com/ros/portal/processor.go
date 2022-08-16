package portal

import (
	"atlas-ros/model"
	"atlas-ros/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func getPortalId(m Model) (uint32, error) {
	return m.Id(), nil
}

func ByNameIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, name string) model.IdProvider[uint32] {
	return func(mapId uint32, name string) model.IdProvider[uint32] {
		return model.ProviderToIdProviderAdapter[Model, uint32](ByNameProvider(l, span)(mapId, name), getPortalId)
	}
}

func RandomProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.Provider[Model] {
	return func(mapId uint32) model.Provider[Model] {
		return model.SliceProviderToProviderAdapter[Model](InMapProvider(l, span)(mapId), model.RandomPreciselyOneFilter[Model])
	}
}

func RandomIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.IdProvider[uint32] {
	return func(mapId uint32) model.IdProvider[uint32] {
		return model.ProviderToIdProviderAdapter[Model, uint32](RandomProvider(l, span)(mapId), getPortalId)
	}
}

func InMapProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.SliceProvider[Model] {
	return func(mapId uint32) model.SliceProvider[Model] {
		return requests.SliceProvider[attributes, Model](l, span)(requestAll(mapId), makePortal)
	}
}

func ByNameProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, portalName string) model.Provider[Model] {
	return func(mapId uint32, portalName string) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestByName(mapId, portalName), makePortal)
	}
}

func makePortal(body requests.DataBody[attributes]) (Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	attr := body.Attributes
	return NewPortalModel(uint32(id), attr.Name, attr.Target, attr.TargetMapId, attr.Type, attr.X, attr.Y, attr.ScriptName), nil
}
