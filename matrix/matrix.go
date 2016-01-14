package matrix

import (
	"bytes"
	"fmt"
)

// An instance of matrix that knows its dimensions and holds a single array
// that represents the data of the matrix.
type Matrix struct {
	n    int         // Row
	m    int         // Columns
	data [][]float32 // size := r * c
}

// Creates a new matrix based on the data provided.
func NewMatrix(d [][]float32) (*Matrix, error) {
	n := len(d)
	sizes := make([]int, n)
	for i, e := range d {
		sizes[i] = len(e)
		if i > 0 && sizes[i-1] != sizes[i] {
			return nil, fmt.Errorf("New matrix data has ragged sizes")
		}
	}

	m := &Matrix{
		n:    n,
		m:    sizes[n-1],
		data: d,
	}

	return m, nil
}

// NewMatrix creates a new matrix with the dimensions (n x m) and filled with zeros.
func EmptyMatrix(n, m int) *Matrix {
	d := make([][]float32, 0)
	for i := 0; i < n; i++ {
		row := make([]float32, m)
		d = append(d, row)
	}

	return &Matrix{
		n:    n,
		m:    m,
		data: d,
	}
}

// Size returns the matrix dimensions Row by Columns (r x c) or (n x m).
func (x *Matrix) Size() (m, n int) {
	return x.n, x.m
}

// inBounds determines if the given i and j are in bounds of the matrix array.
func (x *Matrix) inBounds(i, j int) error {
	if !(0 <= i && i < x.n) {
		return fmt.Errorf("i is out of range: %d (%d x %d)", i, x.m, x.n)
	}
	if !(0 <= j && j < x.m) {
		return fmt.Errorf("j is out of range: %d (%d x %d)", j, x.m, x.n)
	}
	return nil
}

// Item retreives the ith and jth item from the matrix.
func (m *Matrix) Item(i, j int) (float32, error) {
	if err := m.inBounds(i, j); err != nil {
		return 0.0, err
	}
	return m.data[i][j], nil
}

// Set assigns the value k to the ith and jth offset
func (m *Matrix) Set(i, j int, k float32) (*Matrix, error) {
	if err := m.inBounds(i, j); err != nil {
		return m, err
	}
	m.data[i][j] = k
	return m, nil
}

// String provides a pretty string of the values in the matrix in standard shape.
func (x *Matrix) String() string {
	buf := bytes.NewBuffer([]byte{})
	m, n := x.Size()

	buf.WriteString("[")
	k := 0
	for i := 0; i < n; i++ {
		buf.WriteString("\n")
		for j := 0; j < m; j++ {
			f, err := x.Item(i, j)
			if err != nil {
				return err.Error()
			}
			buf.WriteString(fmt.Sprintf("  %f ", f))
			k++
		}
	}
	buf.WriteString("\n]")

	return buf.String()
}
