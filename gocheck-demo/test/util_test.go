package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	Convey("info", t, func() {
		So(Add(1, 1), ShouldEqual, 2)
	})
}
