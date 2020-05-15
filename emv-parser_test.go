package emv

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {

	Convey("Given a EMV Payload", t, func() {

		payload := "9F2606898989898989"
		convey.Printf("Payload: %s", payload)

		Convey("Should build a EMV object from payload", func() {

			emv := NewEMV().Build()

			emv.Parse(payload)

			convey.Printf("TLV: %s", emv.ToTLV())

			So(len(emv.GetEMV().items), ShouldEqual, 1)
			So(emv.ToTLV(), ShouldEqual, payload)

		})
	})
}
