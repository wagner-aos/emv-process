package emv

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {

	Convey("Given a EMV Payload", t, func() {

		Convey("Should build a EMV object", func() {

			emv := NewEMV().Build()

			emv.Parse("9F2606898989898989")

			convey.Printf("TLV: %s", emv.ToBerTLV())

		})
	})
}
