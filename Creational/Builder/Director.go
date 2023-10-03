package Builder

type Director struct {
}

func (d *Director) makeSUV(builder Builder) {
	builder.reset()
	builder.setSeats(4)
	builder.setEngine("SUV engine")
	builder.setGPS()
}

func (d *Director) makeSportCar(builder Builder) {
	builder.reset()
	builder.setSeats(2)
	builder.setEngine("Sport engine")
	builder.setTripComputer()
}
