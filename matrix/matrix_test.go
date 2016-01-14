package matrix

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMatrix(t *testing.T) {

	Convey("Creating a new matrix with ragged data should fail.", t, func() {
		f := [][]float32{
			{1.0, 0.0, 0.0},
			{0.0, 1.0, 0.0},
			{0.0, 0.0, 1.0},
		}
		m, err := NewMatrix(f)
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)

		r, c := m.Size()
		So(r, ShouldEqual, 3)
		So(c, ShouldEqual, 3)
	})

	Convey("Creating a new matrix with data shouldn't fail.", t, func() {
		f := [][]float32{
			{1.0, 0.0, 0.0},
			{0.0, 1.0, 0.0},
			{0.0, 0.0, 1.0},
		}
		m, err := NewMatrix(f)
		So(m, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Creating a new matrix shouldn't produce an error", t, func() {
		x := EmptyMatrix(3, 4)
		m, n := x.Size()
		So(m, ShouldEqual, 3)
		So(n, ShouldEqual, 4)
	})

	Convey("New matrix should have all zero items", t, func() {
		x := EmptyMatrix(3, 4)
		m, n := x.Size()
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				f, err := x.Item(i, j)
				So(err, ShouldBeNil)
				So(f, ShouldEqual, 0.0)
			}
		}
	})

	Convey("Accessing items out of range produces an error", t, func() {
		x := EmptyMatrix(3, 4)
		n, m := x.Size()

		fmt.Printf("%d x %d\n", n, m)
		fmt.Println(x)

		f, err := x.Item(4, 5)
		So(err, ShouldNotBeNil)
		So(f, ShouldEqual, 0.0)
	})
}
