package ch33_test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"

	"github.com/wxb/goLab/geek/goCaiChao/ch33"
)

func TestNewObjPool(t *testing.T) {

	pool := ch33.NewObjPool(10)
	err := pool.Release(&ch33.ReusableObj{})
	assert.Nil(t, err)

	for i := 0; i < 11; i++ {
		obj, err := pool.Obj(1 * time.Second)
		if assert.Nil(t, err) {
			t.Logf("%T\n", obj)
			err = pool.Release(obj)
			assert.Nil(t, err)
		}
	}

	t.Log("Done")
}

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1
		Convey("Important stuff", func() { // This func() will not be executed!
			So("asdf", ShouldEqual, "asdf")
			Convey("More important stuff", func() {
				So("asdf", ShouldEqual, "asdf")

				Convey("More important stuff", func() {
					SkipSo("asdf", ShouldEqual, "asdfe")
					Convey("More important stuff nil", nil)
				})
			})
		})

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 3)
			})
		})
	})
}
