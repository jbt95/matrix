package matrix

import (
	"fmt"
	"log"
	"sync"
)

type Matrix struct {
	data []float64
	rows int
	cols int
}

//Returns a matrix from the given slice or an error if the number of cols and rows does not match the len of the slice
func NewFromSlice(cols int, rows int, data []float64) Matrix {
	if len(data) != rows*cols {
		log.Fatalf("The size of the matrix rows*cols: (%v) doesn't match the size of the given slice (%v)", rows*cols, len(data))
	}
	m := Matrix{cols: cols, rows: rows, data: make([]float64, rows*cols)}
	copy(m.data, data)
	return m
}

//Returns an empty matrix
func NewZeroMatrix(cols, rows int) Matrix {
	m := Matrix{cols: cols, rows: rows, data: make([]float64, rows*cols)}
	for i := range m.data {
		m.data[i] = 0
	}
	return m
}

func Dot(a, b Matrix) Matrix {
	if a.cols != b.rows {
		log.Fatalf("The number of A cols (%v) doesn't match the number of B cols (%v)", a.cols, b.cols)
	}
	var wg sync.WaitGroup
	c := Matrix{rows: a.rows, cols: b.cols, data: make([]float64, a.rows*b.cols)}
	for i := 0; i < a.rows; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < b.cols; j++ {
				for k := 0; k < c.cols; k++ {
					c.data[i*c.cols+k] += a.data[i*a.cols+j] * b.data[j*b.cols+k]
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return c
}

func (m Matrix) Scalar(a float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.data[i*m.cols+j] *= a
		}
	}
}

func (m Matrix) Show() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Print(m.data[i*m.cols+j], " ")
		}
		fmt.Println()
	}
}

func (m Matrix) GetRows() int { return m.rows }
func (m Matrix) GetCols() int { return m.cols }
