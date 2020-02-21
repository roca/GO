package abstract_factory

import (
	"errors"
	"fmt"
)

type IVehicleFactory interface {
	GetVehicle(v int) (IVehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (IVehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with %d not recognized\n", f))
	}
}
