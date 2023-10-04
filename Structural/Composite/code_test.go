package Composite

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestComposite(t *testing.T) {

	convey.Convey("when give a Box", t, func() {
		P1 := &Product{value: 9.9}
		P2 := &Product{value: 13.9}
		P3 := &Product{value: 6.00}
		C2 := &Box{value: 0.1, components: []Component{P2, P3}}
		C1 := &Box{value: 0.2, components: []Component{P1, C2}}
		convey.So(C1.CalculateValue(), convey.ShouldEqual, 30.1)
	})
}
