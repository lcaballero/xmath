package nums

import "fmt"

// Most numbers will be Real numbers or composed of real numbers, but
// so the Real flag will be set, but other flags might also be set.
// These flags do not denote disjoint sets.  A Counting number is
// a WholeNumber, an Integer, and also Real.
const (
	Real = 1 << iota
	Rational
	Irrational
	Integer
	Fraction
	WholeNumber         // Positive integers starting with 0
	WholeNegativeNumber // Integers that are negative
	CountingNumber      // Positive integers starting with 1
	Zero
)

func Dump() {
	fmt.Println(
		Real,
		Rational,
		Irrational,
		Integer,
		Fraction,
		WholeNumber,
		WholeNegativeNumber,
		CountingNumber,
		Zero,
	)
}
