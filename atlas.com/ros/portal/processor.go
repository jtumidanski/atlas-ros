package portal

import (
	"github.com/sirupsen/logrus"
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
			p, err := GetByName(mapId, name)
			if err != nil {
				l.WithError(err).Errorf("Unable to retrieve portal for map %d of name %s. Defaulting to 0.", mapId, name)
				return 0
			}
			return p.Id()
		}
	}
}

func GetByName(mapId uint32, portalName string) (*Model, error) {
	resp, err := requestByName(mapId, portalName)
	if err != nil {
		return nil, err
	}

	p, err := makePortal(resp.Data())
	if err != nil {
		return nil, err
	}
	return p, nil
}

func makePortal(body *dataBody) (*Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	attr := body.Attributes
	return NewPortalModel(uint32(id), attr.Name, attr.Target, attr.TargetMapId, attr.Type, attr.X, attr.Y, attr.ScriptName), nil
}
