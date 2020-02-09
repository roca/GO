package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector interface {
	Construct()
	SetBuilder(b BuildProcess)
}

//Director using the Singleton Pattern
type manufacturingDirector struct {
	builder BuildProcess
}

func (f *manufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}
func (f *manufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

var instance *manufacturingDirector

func GetManufacturerDirector() ManufacturingDirector {
	if instance == nil {
		instance = new(manufacturingDirector) // new() returns a pointer
	}
	return instance
}

//Product
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//A Builder of type car
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

//A Builder of type motorbike
type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

//A Builder of type motorbike
type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 8
	return b
}
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
