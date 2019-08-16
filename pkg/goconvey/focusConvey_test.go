package goconvey_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFocusConvey(t *testing.T) {

	Convey("FocusConvey", t, func() {
		FocusConvey("A", func() {
			// B will not be run
			Convey("B", nil)
			FocusConvey("C", func() {
				// Only D will be run.
				FocusConvey("D", func() {
				})
			})
		})
	})
}

func TestFocusConveyB(t *testing.T) {
	Convey("TestFocusConveyB", t, func() {

		Convey("E", func() {

		})

		Convey("A", func() {
			FocusConvey("C", func() {
				// test B will still run because test D is not marked with Focus
				Convey("B", nil)
				// Mark test D with Focus to run only test D
				FocusConvey("D", func() {
				})
			})
		})
	})
}
