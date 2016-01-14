package prop

import (
	"testing"

	"math"

	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

const fuzz = 0.0001

func Fuzz(val, c float64) {
	So(val, ShouldBeBetweenOrEqual, c-fuzz, c+fuzz)
}

func data1() []float64 {
	return []float64{
		0.684, 2.540, 0.924, 3.130, 1.038, 0.598,
		0.483, 3.520, 1.285, 2.650, 1.497,
	}
}

func TestProb(t *testing.T) {

	Convey("UpperFourth() cases", t, func() {
		cases := []struct {
			size int
			a int
			b int
			err error
		}{
			{ 11, 8, 8, nil },
			{  4, 2, 3, nil },
			{  3, 2, 2, nil },
			{  2, 1, 1, nil },
			{  1, 0, 0, ValueDoesNotExist },
			{  0, 0, 0, ValueDoesNotExist },
			{ -1, 0, 0, ValueDoesNotExist },
		}

		for _, expected := range cases {
			a, b, err := UpperFourth(expected.size)
			c := expected
			fmt.Println(c.size, c.a, c.b)
			So(a, ShouldEqual, expected.a)
			So(b, ShouldEqual, expected.b)
			So(err, ShouldEqual, expected.err)
		}
	})

	Convey("LowerFourth() cases", t, func() {
		cases := []struct {
			size int
			a int
			b int
			err error
		}{
			{ 11, 2, 2, nil },
			{  4, 0, 1, nil },
			{  3, 0, 0, nil },
			{  2, 0, 0, nil },
			{  1, 0, 0, ValueDoesNotExist },
			{  0, 0, 0, ValueDoesNotExist },
			{ -1, 0, 0, ValueDoesNotExist },
		}

		for _, expected := range cases {
			a, b, err := LowerFourth(expected.size)
			c := expected
			fmt.Println(c.size, c.a, c.b)
			So(a, ShouldEqual, expected.a)
			So(b, ShouldEqual, expected.b)
			So(err, ShouldEqual, expected.err)
		}
	})

	Convey("Middle() cases", t, func() {
		cases := []struct {
			size int
			a int
			b int
			err error
		}{
			{  4, 1, 2, nil },
			{  3, 1, 1, nil },
			{  2, 0, 1, nil },
			{  1, 0, 0, nil },
			{  0, 0, 0, ValueDoesNotExist },
			{ -1, 0, 0, ValueDoesNotExist },
		}

		for _, expected := range cases {
			a, b, err := Middle(expected.size)
			So(a, ShouldEqual, expected.a)
			So(b, ShouldEqual, expected.b)
			So(err, ShouldEqual, expected.err)
		}
	})

	Convey("MiddleIndex for empty set should produce an error.", t, func() {
		vals := []float64{}
		_, _, err := Middle(len(vals))
		So(err, ShouldNotBeNil)
	})

	Convey("Variance.", t, func() {
		nums := data1()
		v := StdDeviation(nums)
		var sumSqrs float64 = 11.9359
		var variance float64 = sumSqrs / float64(len(nums)-1)
		Fuzz(v, math.Sqrt(variance))
	})

	Convey("Mean of data should produce correct value.", t, func() {
		sum := Mean(data1())
		Fuzz(sum, 1.6681)
	})

	Convey("Finds the middle value for empty array.", t, func() {
		nums := []float64{}
		med := Median(nums, false)
		Fuzz(med, 0.0)
	})

	Convey("Finds the middle value for nil array.", t, func() {
		var nums []float64 = nil
		med := Median(nums, false)
		Fuzz(med, 0.0)
	})

	Convey("Finds the middle value for small arrays.", t, func() {
		nums := []float64{2.3, -1.1}
		med := Median(nums, false)
		c := (2.3 - 1.1) / 2.0
		Fuzz(med, c)

		nums = []float64{1.4}
		med = Median(nums, false)
		c = 1.4
		Fuzz(med, c)
	})

	Convey("Finds the middle number (median) of the values.", t, func() {
		nums := []float64{1.2, 2.3, -1.1, -10.1, 9.1}
		med := Median(nums, false)
		Fuzz(med, 1.2)

		nums = []float64{1.2, 2.3, -1.1, -10.1, 9.1, 1.4}
		med = Median(nums, false)
		Fuzz(med, 1.3)
	})

	Convey("Combination(size, draw) of 8, 2 should equal 28 ", t, func() {
		total := Combination(2, 8)
		var a, b, k int64 = 8, 7, 2

		So(total, ShouldEqual, (a*b)/k)
	})

	Convey("Permutation with `draw` from `size` === size! / (size - draw)! ", t, func() {
		total := Permutation(2, 8)
		var a, b int64 = 8, 7
		So(total, ShouldEqual, a*b)
	})

	Convey("Factorial(4) should be 2 x 3 x 4", t, func() {
		total := Factorial(4)
		var a, b, c int64 = 2, 3, 4
		So(total, ShouldEqual, a*b*c)
	})

	Convey("Sorting check", t, func() {
		nums := []float64{1.2, 2.3, -1.1, -10.1, 10.1}
		Sort(nums)

		So(nums[0], ShouldEqual, -10.1)
		So(nums[1], ShouldEqual, -1.1)
		So(nums[2], ShouldEqual, 1.2)
		So(nums[3], ShouldEqual, 2.3)
		So(nums[4], ShouldEqual, 10.1)
	})
}
