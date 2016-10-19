// Package gfmatrix implements basic operations on matrices over Rijndael's field and the random generation of new ones.
package gfmatrix

import (
	"fmt"
)

// Matrix represents a GF(2^8)-matrix.
type Matrix []Row

// Mul right-multiplies a matrix by a row.
func (e Matrix) Mul(f Row) Row {
	out, in := e.Size()
	if in != f.Size() {
		panic("Can't multiply by row that is wrong size!")
	}

	res := NewRow(out)
	for i := 0; i < out; i++ {
		res[i] = e[i].DotProduct(f)
	}

	return res
}

// Add adds two matrices from GF(2^8)^nxm.
func (e Matrix) Add(f Matrix) Matrix {
	a, _ := e.Size()

	out := make([]Row, a)
	for i, _ := range out {
		out[i] = e[i].Add(f[i])
	}

	return out
}

// Compose returns the result of composing e with f.
func (e Matrix) Compose(f Matrix) Matrix {
	n, m := e.Size()
	p, q := f.Size()

	if m != p {
		panic("Can't multiply matrices of the wrong size!")
	}

	out := GenerateEmpty(n, q)
	g := f.Transpose()

	for i, e_i := range e {
		for j, g_j := range g {
			out[i][j] = e_i.DotProduct(g_j)
		}
	}

	return out
}

// Transpose returns the transpose of a matrix.
func (e Matrix) Transpose() Matrix {
	n, m := e.Size()
	out := GenerateEmpty(m, n)

	for i, row := range e {
		for j, elem := range row {
			out[j][i] = elem.Dup()
		}
	}

	return out
}

// Invert computes the multiplicative inverse of a matrix, if it exists.
func (e Matrix) Invert() (Matrix, bool) {
	inv, _, frees := e.gaussJordan()
	return inv, len(frees) == 0
}

// FindPivot finds a row with non-zero entry in column col, starting at the given row and moving down. It returns the
// index of the row or -1 if one does not exist.
func (e Matrix) FindPivot(row, col int) int {
	out, _ := e.Size()

	for i := row; i < out; i++ {
		if !e[i][col].IsZero() {
			return i
		}
	}

	return -1
}

// Dup returns a duplicate of this matrix.
func (e Matrix) Dup() Matrix {
	n, m := e.Size()
	out := GenerateEmpty(n, m)

	for i, row := range e {
		for j, elem := range row {
			out[i][j] = elem.Dup()
		}
	}

	return out
}

// IsBinary returns true if the matrix contains only zero and one entries.
func (e Matrix) IsBinary() bool {
	for _, row := range e {
		for _, col := range row {
			if !col.IsZero() && !col.IsOne() {
				return false
			}
		}
	}

	return true
}

// Equals returns true if two matrices are equal and false otherwise.
func (e Matrix) Equals(f Matrix) bool {
	a, _ := e.Size()
	b, _ := f.Size()

	if a != b {
		return false
	}

	for row := 0; row < a; row++ {
		if !e[row].Equals(f[row]) {
			return false
		}
	}

	return true
}

// Size returns the dimensions of the matrix in (Rows, Columns) order.
func (e Matrix) Size() (int, int) {
	if len(e) == 0 {
		return 0, 0
	} else {
		return len(e), e[0].Size()
	}
}

func (e Matrix) String() string {
	out := []rune{}

	for _, row := range e {
		out = append(out, []rune(row.String())...)
		out = append(out, '\n')
	}

	return string(out)
}

// OctaveString converts the matrix into a string that can be imported into Octave.
func (e Matrix) OctaveString() string {
	out := []rune{}

	for _, row := range e {
		out = append(out, []rune(row.OctaveString())...)
	}

	return string(out)
}

func (e Matrix) GoString() string {
	out := []rune("gfmatrix.Matrix{\n")

	for _, row := range e {
		out = append(out, []rune("\tgfmatrix.Row{")...)

		for _, elem := range row[:len(row)-1] {
			out = append(out, []rune(fmt.Sprintf("0x%2.2x, ", elem))...)
		}

		out = append(out, []rune(fmt.Sprintf("0x%2.2x},\n", row[len(row)-1]))...)
	}

	out = append(out, '}')
	return string(out)
}
