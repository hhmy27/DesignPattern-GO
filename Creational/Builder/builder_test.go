package Builder

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBuilder(t *testing.T) {
	var d *Director
	Convey("When give a builder", t, func() {
		Convey("When make a sport car", func() {
			builder := &CarBuilder{}
			d.makeSportCar(builder)
			result := builder.getResult()
			So(result, ShouldHaveSameTypeAs, &Car{})
			So(result.Seats, ShouldEqual, 2)
			So(result.Engine, ShouldEqual, "Sport engine")
			So(result.Computer, ShouldEqual, true)
		})
		Convey("When make a suv car manual", func() {
			builder := &CarManualBuilder{}
			d.makeSUV(builder)
			result := builder.getResult()
			So(result, ShouldHaveSameTypeAs, &CarManual{})
			So(result.Seats, ShouldEqual, 4)
			So(result.Engine, ShouldEqual, "SUV engine")
			So(result.GPS, ShouldEqual, true)
			So(result.Computer, ShouldEqual, false)
		})
	})
}
