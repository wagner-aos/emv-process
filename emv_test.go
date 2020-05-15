package emv

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEMV(t *testing.T) {

	Convey("Given a Tag and EMV Builder", t, func() {

		Convey("Should build a emv tag", func() {

			tag := New().
				SetTag("5F20", "WAGNER ALVES").
				BuildTag()

			convey.Printf("%v", tag)

			So(tag.GetName(), ShouldEqual, "5F20")
			So(tag.GetValue(), ShouldEqual, "WAGNER ALVES")
		})

		Convey("Should build a emv tag list and add tag", func() {

			emv := NewEMV().
				AddTag("5F20", "WAGNER ALVES").
				AddTag("9F0B", "WAGNER ALVES DE OLIVEIRA SILVA").
				Build()

			for _, t := range emv.GetEMV().items {
				convey.Printf("\n\tTAG: %s: %s", t.GetName(), t.GetValue())
			}

			So(len(emv.GetEMV().items), ShouldEqual, 2)

		})

		Convey("Should build a emv tag list and remove a tag", func() {

			emv := NewEMV().
				AddTag("5F20", "WAGNER ALVES").
				AddTag("9F0B", "WAGNER ALVES DE OLIVEIRA SILVA").
				Build()

			emv.RemoveTag("9F0B")

			for _, t := range emv.GetEMV().items {
				convey.Printf("\n\tTAG: %s: %s", t.GetName(), t.GetValue())
			}

			So(len(emv.GetEMV().items), ShouldEqual, 1)

		})

		Convey("Should build a emv tag list and get a tag by name", func() {

			emv := NewEMV().
				AddTag("5F20", "WAGNER ALVES").
				AddTag("9F0B", "WAGNER ALVES DE OLIVEIRA SILVA").
				Build()

			tagName, tagValue, tagSize, err := emv.GetTag("5F20")
			convey.Printf("\n\tTag Name: %s", tagName)
			convey.Printf("\n\tTag Value: %s", tagValue)
			convey.Printf("\n\tTag Size: %d", tagSize)

			So(tagName, ShouldEqual, "5F20")
			So(tagValue, ShouldEqual, "WAGNER ALVES")
			So(tagSize, ShouldEqual, 12)
			So(err, ShouldBeNil)

			So(len(emv.GetEMV().items), ShouldEqual, 2)

		})

		Convey("Should return an error if a tag not exists", func() {

			emv := NewEMV().
				AddTag("5F20", "WAGNER ALVES").
				AddTag("9F0B", "WAGNER ALVES DE OLIVEIRA SILVA").
				Build()

			tagName, tagValue, tagSize, err := emv.GetTag("XXXX")

			So(tagName, ShouldBeEmpty)
			So(tagValue, ShouldBeEmpty)
			So(tagSize, ShouldEqual, 0)
			So(err, ShouldBeError)

			So(len(emv.GetEMV().items), ShouldEqual, 2)

		})

	})

}

func TestEMVParser(t *testing.T) {

	Convey("Should build a emv tag list and generate TLV", t, func() {

		emv := NewEMV().
			AddTag("82", "5800").
			AddTag("84", "A0000000041010").
			AddTag("95", "4200008000").
			AddTag("9F26", "54BD15AE210A5117").
			AddTag("9F27", "80").
			AddTag("9F10", "0210A50002020000000000000000000000FF").
			AddTag("9F36", "00F6").
			AddTag("9F37", "B576D113").
			AddTag("9A", "190322").
			AddTag("9C", "00").
			AddTag("9F02", "000000025800").
			AddTag("9F03", "000000000000").
			AddTag("5F2A", "0986").
			AddTag("9F1A", "0076").
			AddTag("9F35", "22").
			AddTag("9F34", "410302").
			AddTag("5F24", "220531").
			AddTag("9F33", "E0F0C8").
			AddTag("5F28", "0076").
			// AddTag("5F20", "WAGNER ALVES").
			// AddTag("9F0B", "WAGNER ALVES DE OLIVEIRA SILVA").
			Build()

		for _, t := range emv.GetEMV().items {
			convey.Printf("\n\tTAG: %s: %s", t.GetName(), t.GetValue())
		}

		tlv := emv.ToTLV()
		convey.Printf("\n\tTLV: %s", tlv)

		//So(len(emv.GetEMV().items), ShouldEqual, 21)

	})
}
