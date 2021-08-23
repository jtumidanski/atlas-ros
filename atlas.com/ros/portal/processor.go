package portal

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
)

type IdProvider func() uint32

func FixedPortalIdProvider(portalId uint32) IdProvider {
	return func() uint32 {
		return portalId
	}
}

func ByNamePortalIdProvider(l logrus.FieldLogger) func(mapId uint32, name string) IdProvider {
	return func(mapId uint32, name string) IdProvider {
		return func() uint32 {
			p, err := GetByName(l)(mapId, name)
			if err != nil {
				l.WithError(err).Errorf("Unable to retrieve portal for map %d of name %s. Defaulting to 0.", mapId, name)
				return 0
			}
			return p.Id()
		}
	}
}

func RandomPortalIdProvider(l logrus.FieldLogger) func(mapId uint32) IdProvider {
	return func(mapId uint32) IdProvider {
		return func() uint32 {
			ps, err := ForMap(l)(mapId)
			if err != nil {
				l.WithError(err).Errorf("Unable to retrieve portals for map %d. Defaulting to 0.", mapId)
				return 0
			}
			if len(ps) == 0 {
				l.Warnf("No portals in map %d. Defaulting to zero.", mapId)
				return 0
			}
			return ps[rand.Intn(len(ps))].Id()
		}
	}
}

func ForMap(l logrus.FieldLogger) func(mapId uint32) ([]*Model, error) {
	return func(mapId uint32) ([]*Model, error) {
		resp, err := requestAll(l)(mapId)
		if err != nil {
			return nil, err
		}

		results := make([]*Model, 0)
		for _, d := range resp.DataList() {
			p, err := makePortal(d)
			if err != nil {
				return nil, err
			}
			results = append(results, p)
		}
		return results, nil
	}
}

func GetByName(l logrus.FieldLogger) func(mapId uint32, portalName string) (*Model, error) {
	return func(mapId uint32, portalName string) (*Model, error) {
		resp, err := requestByName(l)(mapId, portalName)
		if err != nil {
			return nil, err
		}

		p, err := makePortal(resp.Data())
		if err != nil {
			return nil, err
		}
		return p, nil
	}
}

func makePortal(body *dataBody) (*Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	attr := body.Attributes
	return NewPortalModel(uint32(id), attr.Name, attr.Target, attr.TargetMapId, attr.Type, attr.X, attr.Y, attr.ScriptName), nil
}
