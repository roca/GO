package builder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := GetManufacturerDirector()

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	assert.Equal(t, car.Wheels, 4, fmt.Sprintf("A Car has 4 wheels not %d", car.Wheels))
	assert.Equal(t, car.Structure, "Car", "A %s is not a car", car.Structure)
	assert.Equal(t, car.Seats, 5, "This Car should have 5 not %s", car.Seats)

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()

	assert.Equal(t, motorbike.Wheels, 2, fmt.Sprintf("A MotorBike has 2 wheels not %d", motorbike.Wheels))
	assert.Equal(t, motorbike.Structure, "Motorbike", "A %s is not a MotorBike", motorbike.Structure)
	assert.Equal(t, motorbike.Seats, 2, "This MotorBike should have 1 not %s", motorbike.Seats)
	motorbike.Seats = 1
	assert.Equal(t, motorbike.Seats, 1, "This MotorBike Seats can be changed to 1 it's still %s", motorbike.Seats)

	busBuilder := &BusBuilder{}
	manufacturingComplex.SetBuilder(busBuilder)
	manufacturingComplex.Construct()

	bus := busBuilder.GetVehicle()

	assert.Equal(t, bus.Wheels, 8, fmt.Sprintf("A Bus has 8 wheels not %d", bus.Wheels))
	assert.Equal(t, bus.Structure, "Bus", "A %s is not a Bus", bus.Structure)
	assert.Equal(t, bus.Seats, 30, "This MotorBike should have 1 not %s", bus.Seats)

}
