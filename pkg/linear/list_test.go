package linear_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yao-sh/goext/pkg/linear"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Adding element to Linked List", t, func() {
		ll := linear.LinkedList{}
		Convey("When the integer is incremented", func() {
			ll.Add(1)
			Convey("The length of the linked list should equal to one", func() {
				So(ll.Length(), ShouldEqual, 1)
			})
		})
		ll.Display()
	})
}
