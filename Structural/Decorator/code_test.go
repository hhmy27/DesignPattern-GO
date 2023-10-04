package Decorator

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTomatoTopping(t *testing.T) {
	convey.Convey("when add cheese topping", t, func() {
		v := &VeggieMania{}
		c := &CheeseTopping{pizza: v}
		convey.So(c.getPrice(), convey.ShouldEqual, 25)
	})
}

func TestCheeseTopping(t *testing.T) {
	convey.Convey("when add cheese topping", t, func() {
		v := &VeggieMania{}
		t := &TomatoTopping{pizza: v}
		convey.So(t.getPrice(), convey.ShouldEqual, 22)
	})
}

func TestTomatoAndCheeseTopping(t *testing.T) {
	convey.Convey("when give a VeggieMania", t, func() {
		v := &VeggieMania{}
		convey.Convey("when give a cheese and tomato topping", func() {
			c := &CheeseTopping{pizza: v}
			convey.So(c.getPrice(), convey.ShouldEqual, 25)
			t := &TomatoTopping{pizza: c}
			convey.So(t.getPrice(), convey.ShouldEqual, 32)
		})

	})
}
