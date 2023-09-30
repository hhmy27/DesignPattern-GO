package AbstractFactory

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	Convey("Given a ModernCreator", t, func() {
		var creator ICreator
		creator = &ModernCreator{}

		Convey("When calling getTable method", func() {
			table := creator.getTable()

			Convey("It should return a ModernTable", func() {
				So(table, ShouldHaveSameTypeAs, &ModernTable{})
			})
		})

		Convey("When calling getSeat method", func() {
			seat := creator.getSeat()

			Convey("It should return a ModernSeat", func() {
				So(seat, ShouldHaveSameTypeAs, &ModernSeat{})
			})
		})
	})

	Convey("Given an ArtCreator", t, func() {
		var creator ICreator    // 声明一个Creator接口类型的变量
		creator = &ArtCreator{} // 将具体的ArtCreator绑定到Creator接口变量

		Convey("When calling getTable method", func() {
			table := creator.getTable()

			Convey("It should return an ArtTable", func() {
				So(table, ShouldHaveSameTypeAs, &ArtTable{})
			})
		})

		Convey("When calling getSeat method", func() {
			seat := creator.getSeat()

			Convey("It should return an ArtSeat", func() {
				So(seat, ShouldHaveSameTypeAs, &ArtSeat{})
			})
		})
	})
}
