package matrix

import (
	"fmt"
)

//row*ncols + col
type Matrix struct {
	data []float64
	rows int
	cols int
	dim  int
}

//Returns a matrix from the given slice or an error if the number of cols and rows does not match the len of the slice
func NewFromSlice(cols int, rows int, data []float64) (*Matrix, error) {
	if len(data) != rows*cols {
		return nil, fmt.Errorf("The size of the matrix rows*cols: (%v) doesn't match the size of the given slice (%v)", rows*cols, len(data))
	}
	m := &Matrix{cols: cols, rows: rows, dim: rows * cols, data: make([]float64, rows*cols)}
	copy(m.data, data)
	return m, nil
}

//Returns an empty matrix
func NewZeroMatrix(cols, rows int) *Matrix {
	m := &Matrix{cols: cols, rows: rows, dim: rows * cols, data: make([]float64, rows*cols)}
	for i := range m.data {
		m.data[i] = 0
	}
	return m
}

func (m *Matrix) Scalar(a float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.data[i*m.cols+j] *= a
		}
	}
}

func (m *Matrix) Show() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Print(m.data[i*m.cols+j], " ")
		}
		fmt.Println()
	}
}
