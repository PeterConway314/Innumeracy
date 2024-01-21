package pi_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"pi/pi"
)

func TestGenerateXY(t *testing.T) {
	Convey("Given we are testing GenerateXY", t, func() {
		for i := 0; i < 25; i++ {
			x, y := pi.GenerateXY()
			So(x, ShouldBeLessThanOrEqualTo, 1)
			So(y, ShouldBeLessThanOrEqualTo, 1)
		}
	})
}

func TestIsCoordinateInsideUnitCircle(t *testing.T) {
	Convey("Given we are testing IsCoordinateInsideUnitCircle", t, func() {
		type testCase struct {
			x      float64
			y      float64
			inside bool
		}
		cases := []testCase{
			{1, 1, false},
			{0.000001, 1, false},
			{0.1, 0.1, true},
			{1, 0, true},
			{0, 1, true},
		}
		for _, tc := range cases {
			description := "outside"
			if tc.inside {
				description = "inside"
			}
			Convey(fmt.Sprintf("When x=%f and y=%f we should be %s a unit circle", tc.x, tc.y, description), func() {
				So(pi.IsCoordinateInsideUnitCircle(tc.x, tc.y), ShouldEqual, tc.inside)
			})
		}
	})
}
