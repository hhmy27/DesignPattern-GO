package Builder

type Builder interface {
	reset()
	setSeats(number int)
	setEngine(engine string)
	setTripComputer()
	setGPS()
}

type CarBuilder struct {
	Car *Car
}

func (c *CarBuilder) reset() {
	c.Car = &Car{}
}

func (c *CarBuilder) setSeats(number int) {
	c.Car.Seats = number
}

func (c *CarBuilder) setEngine(engine string) {
	c.Car.Engine = engine
}

func (c *CarBuilder) setTripComputer() {
	c.Car.Computer = true
}

func (c *CarBuilder) setGPS() {
	c.Car.GPS = true
}

func (c *CarBuilder) getResult() *Car {
	return c.Car
}

type CarManualBuilder struct {
	CarManual *CarManual
}

func (c *CarManualBuilder) reset() {
	c.CarManual = &CarManual{}
}

func (c *CarManualBuilder) setSeats(number int) {
	c.CarManual.Seats = number
}

func (c *CarManualBuilder) setEngine(engine string) {
	c.CarManual.Engine = engine
}

func (c *CarManualBuilder) setTripComputer() {
	c.CarManual.Computer = true
}

func (c *CarManualBuilder) setGPS() {
	c.CarManual.GPS = true
}

// In some cases, the products you build may not be of the same type.
// Therefore, you need a getResult() function to obtain the result of the build.
func (c *CarManualBuilder) getResult() *CarManual {
	return c.CarManual
}
