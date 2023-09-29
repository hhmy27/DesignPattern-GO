package Factory

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateTable(t *testing.T) {
	Convey("Given a Creator", t, func() {
		var creator ICreator
		Convey("Given a TableCreator", func() {
			creator = &TableCreator{}
			product := creator.createProduct()
			Convey("When calling createProduct method", func() {

				Convey("It should return a Table IProduct", func() {
					So(product, ShouldHaveSameTypeAs, &Table{})
				})

				Convey("And the Table IProduct should be usable", func() {
					tableProduct, ok := product.(*Table)
					So(ok, ShouldBeTrue)
					So(tableProduct.use(), ShouldEqual, "Using Table")
				})
			})
		})
		Convey("Given a SeatCreator", func() {
			creator = &SeatCreator{}
			product := creator.createProduct()
			Convey("When calling createProduct method", func() {

				Convey("It should return a Seat IProduct", func() {
					So(product, ShouldHaveSameTypeAs, &Seat{})
				})

				Convey("And the Seat IProduct should be usable", func() {
					tableProduct, ok := product.(*Seat)
					So(ok, ShouldBeTrue)
					So(tableProduct.use(), ShouldEqual, "Using Seat")
				})
			})
		})
	})

}
