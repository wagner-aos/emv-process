package emv

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEMVTag(t *testing.T) {

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

	})

}
