package emv

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEMVTag(t *testing.T) {

	Convey("Given a data", t, func() {

		Convey("Should builder a emv tag", func() {

			builder := New()
			tag := builder.
				SetName("wagner").
				SetValue("1010").
				Build()

			convey.Printf("%v", tag)

			So(tag.GetName(), ShouldEqual, "wagner")
			So(tag.GetValue(), ShouldEqual, "1010")

		})

	})

}
