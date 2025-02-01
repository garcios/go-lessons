package pkg

import (
	"encoding/json"
	"errors"
)

type Vehicle interface {
	Info()
}

// Validate interface compliance.
var _ Vehicle = (*Car)(nil)
var _ Vehicle = (*Airplane)(nil)

type Fleet struct {
	CityLocation string    `json:"city_location"`
	Vehicles     []Vehicle `json:"vehicles"`
}

func (f *Fleet) Unmarshal(bytes []byte) error {
	fleet := &struct {
		CityLocation string            `json:"city_location"`
		Vehicles     []json.RawMessage `json:"vehicles"`
	}{}

	err := json.Unmarshal(bytes, &fleet)
	if err != nil {
		return err
	}

	f.CityLocation = fleet.CityLocation
	f.Vehicles = make([]Vehicle, len(fleet.Vehicles))

	for i, v := range fleet.Vehicles {
		vehicleType := &struct {
			Type string `json:"type"`
		}{}

		err := json.Unmarshal(v, &vehicleType)
		if err != nil {
			return errors.New("missing Type attribute in vehicle")
		}

		switch vehicleType.Type {
		case "CAR":
			var c Car
			err := json.Unmarshal(v, &c)
			if err != nil {
				return errors.New("failed to unmarshal vehicle as Car")
			}
			f.Vehicles[i] = &c
		case "PLANE":
			var a Airplane
			err := json.Unmarshal(v, &a)
			if err != nil {
				return errors.New("failed to unmarshal vehicle as Airplane")
			}
			f.Vehicles[i] = &a
		default:
			return errors.New("unknown vehicle type")
		}
	}

	return nil
}
