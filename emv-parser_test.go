package emv

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {

	Convey("Given a EMV Payload", t, func() {

		//payload := "9F260854BD15AE210A51179F2701809F10180210A50002020000000000000000000000FF"
		//payload := "9F10180210A50002020000000000000000000000FF"
		payload := "9F260854BD15AE210A51179F2701809F10120210A50002020000000000000000000000FF9F3704B576D1139F360200F6950542000080009A031903229C01009F02060000000258009F03060000000000005F2A020986820258009F1A0200769F3501229F34034103025F24032205319F3303E0F0C85F280200768407A0000000041010"
		convey.Printf("%s", payload)

		Convey("Should build a EMV object from payload", func() {

			emv := NewEMV().Build()

			emv.Parse(payload)

			tlv := emv.ToTLV()
			convey.Printf("\n\tTLV: %s", tlv)

			So(len(emv.GetEMV().items), ShouldEqual, 19)
			So(len(tlv), ShouldEqual, len(payload))

		})
	})
}