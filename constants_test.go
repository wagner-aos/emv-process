package emv

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstants(t *testing.T) {

	Convey("Should validate map of constants", t, func() {

		ok := validateTagNameMap()

		So(ok, ShouldBeTrue)

	})

}
