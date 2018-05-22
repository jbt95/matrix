package matrix

import (
	"fmt"
)

type Row []float64

type Matrix struct {
	P    [][]float64
	Rows int
	Cols int
}

func New(rows, cols int, elements ...Row) (*Matrix, error) {
	if len(elements) != rows {
		return nil, fmt.Errorf("The number of elements passed (%d) doesn't match the number of rows (%d)", len(elements), rows)
	}

	m := &Matrix{
		Cols: cols,
		Rows: rows,
		P:    allocateMatrix(cols, rows),
	}

	for i := range elements {
		if len(elements[i]) != cols {
			return nil, fmt.Errorf("The row number %d doesn't have %d columns", i+1, cols)
		}
		m.P[i] = elements[i]
	}

	return m, nil
}

func (m *Matrix) Scalar(a float64) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.P[i][j] *= a
		}
	}
}

func (m *Matrix) Show() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("%f ", m.P[i][j])
		}
		fmt.Println()
	}
}

func allocateMatrix(cols, rows int) [][]float64 {
	p := make([][]float64, rows)
	for i := 0; i < cols; i++ {
		p[i] = make([]float64, cols)
	}
	return p
}
