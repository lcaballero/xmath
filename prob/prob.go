package prop

import (
	"errors"
	"math"
	"sort"
)


var ValueDoesNotExist = errors.New("Value does not exist")

func Sum(nums []float64) float64 {
	total := 0.0
	for _, f := range nums {
		total += f
	}
	return total
}

func Mean(nums []float64) float64 {
	total := Sum(nums)
	n := float64(len(nums))
	return total / n
}

func VarianceSquared(nums []float64) float64 {
	mean := Mean(nums)
	total := 0.0
	for _, f := range nums {
		a := f - mean
		total += (a * a)
	}
	n := float64(len(nums) - 1)
	variance := total / n
	return variance
}

func StdDeviation(nums []float64) float64 {
	s2 := VarianceSquared(nums)
	return math.Sqrt(s2)
}

func Factorial(n int64) int64 {
	total := n
	for i := total - 1; i > 0; i-- {
		total = total * i
	}
	return total
}

func Permutation(draw, size int64) int64 {
	total := size
	for i := total - 1; i > size-draw; i-- {
		total = total * i
	}
	return total
}

// Combination calculates the combination with the given draw and size.
func Combination(draw, size int64) int64 {
	return Permutation(draw, size) / Factorial(draw)
}

// Sort orders the array.
func Sort(nums []float64) {
	sort.Float64Slice(nums).Sort()
}

// Middle finds two indices which produce the edges of the lower and upper
// half.  When n is odd the indexes will be the same, but when even the first
// value will be one less than the upper edge.  If n <= 0 then an error
// will be returned as well.
func Middle(n int) (int, int, error) {
	if n <= 0 {
		return 0, 0, ValueDoesNotExist
	}
	if n == 1 {
		return 0, 0, nil
	}

	m := n / 2 // Have: n > 2
	even := (n % 2) == 0

	if even {
		return m - 1, m, nil
	} else {
		return m, m, nil
	}
}

// LowerFourth produces the index for the lower fourth.
func LowerFourth(n int) (int, int, error) {
	if n < 2 {
		return 0, 0, ValueDoesNotExist
	}
	_, b, err := Middle(n)
	if err != nil {
		return 0, 0, err
	}
	return Middle(b)
}

func UpperFourth(n int) (int, int, error) {
	if n < 2 {
		return 0, 0, ValueDoesNotExist
	}
	k, _, err := Middle(n)
	if err != nil {
		return 0, 0, err
	}
	a, b, err := Middle(n-k+1)
	if err != nil {
		return 0, 0, err
	}
	return a+k, b+k, nil
}

// Median sorts nums if isSorted is false, and then returns the middle value.
func Median(nums []float64, isSorted bool) float64 {
	if nums == nil || len(nums) == 0 {
		return 0.0
	}
	if !isSorted {
		Sort(nums)
	}

	a, b, err := Middle(len(nums))

	if err != nil && a == b {
		return nums[a]
	} else {
		return (nums[a] + nums[b]) / 2.0
	}
}
