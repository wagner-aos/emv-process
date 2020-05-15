package cpm

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEMVTag(t *testing.T) {

	Convey("Given a Tag and EMV Builder", t, func() {

		Convey("Should build a emv tag", func() {
			emv := new(EMVQR)

			cdt := new(CommonDataTemplate)
			cdt.DataApplicationPAN = "1234567890123458"
			cdt.DataCardholderName = "WAGNER ALVES"
			cdt.DataLanguagePreference = "ruesdeen"

			cdtt := new(CommonDataTransparentTemplate)
			cdtt.DataIssuerApplicationData = "06010A03000000"
			cdtt.DataApplicationCryptogram = "584FD385FA234BCC"
			cdtt.DataApplicationTransactionCounter = "0001"
			cdtt.DataUnpredictableNumber = "6D58EF13"
			cdt.CommonDataTransparentTemplates = append(cdt.CommonDataTransparentTemplates, *cdtt)

			emv.CommonDataTemplates = append(emv.CommonDataTemplates, *cdt)

			emvCode, err := emv.GeneratePayload()

			convey.Printf("%s", emvCode)

			//So(emvCode, ShouldEqual, "5F20")
			So(err, ShouldBeNil)

			// hQVDUFYwMWETTwegAAAAVVVVUAhQcm9kdWN0MWETTwegAAAAZmZmUAhQcm9kdWN0MmJJWggSNFZ4kBI0WF8gDkNBUkRIT0xERVIvRU1WXy0IcnVlc2RlZW5kIZ8QBwYBCgMAAACfJghYT9OF+iNLzJ82AgABnzcEbVjvEw==
		})

	})
}
