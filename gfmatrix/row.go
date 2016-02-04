package gfmatrix

import (
	"fmt"

	"github.com/OpenWhiteBox/primitives/number"
)

// Row is a row / vector of elements from GF(2^8).
type Row []number.ByteFieldElem

// NewRow returns an empty n-component row.
func NewRow(n int) Row {
	return Row(make([]number.ByteFieldElem, n))
}

// LessThan returns true if row i is "less than" row j. If you use sort a permutation matrix according to LessThan,
// you'll always get the identity matrix.
func LessThan(i, j Row) bool {
	if i.Size() != j.Size() {
		panic("Can't compare rows that are different sizes!")
	}

	for k, _ := range i {
		if i[k] != 0x00 || j[k] != 0x00 {
			if j[k] == 0x00 {
				return true
			} else {
				return false
			}
		}
	}

	return false
}

// Add adds two vectors from GF(2^8)^n.
func (e Row) Add(f Row) Row {
	if e.Size() != f.Size() {
		panic("Can't add rows that are different sizes!")
	}

	out := e.Dup()
	for i, f_i := range f {
		out[i] = out[i].Add(f_i)
	}

	return out
}

// ScalarMul multiplies a row by a scalar.
func (e Row) ScalarMul(f number.ByteFieldElem) Row {
	out := e.Dup()
	for i, _ := range out {
		out[i] = out[i].Mul(f)
	}

	return out
}

// DotProduct computes the dot product of two vectors.
func (e Row) DotProduct(f Row) number.ByteFieldElem {
	if e.Size() != f.Size() {
		panic("Can't compute dot product of two vectors of different sizes!")
	}

	res := number.ByteFieldElem(0x00)
	for i, _ := range e {
		res = res.Add(e[i].Mul(f[i]))
	}

	return res
}

// IsPermutation returns true if the row is a permutation of the first len(e) elements of GF(2^8) and false otherwise.
func (e Row) IsPermutation() bool {
	sums := [256]int{}
	for _, e_i := range e {
		sums[e_i]++
	}

	for _, x := range sums[0:len(e)] {
		if x != 1 {
			return false
		}
	}

	return true
}

// Height returns the position of the first non-zero entry in the row, or -1 if the row is zero.
func (e Row) Height() int {
	for i, e_i := range e {
		if !e_i.IsZero() {
			return i
		}
	}

	return -1
}

// Equals returns true if two rows are equal and false otherwise.
func (e Row) Equals(f Row) bool {
	if e.Size() != f.Size() {
		panic("Can't compare rows that are different sizes!")
	}

	for i, _ := range e {
		if e[i] != f[i] {
			return false
		}
	}

	return true
}

// IsZero returns whether or not the row is identically zero.
func (e Row) IsZero() bool {
	for _, e_i := range e {
		if !e_i.IsZero() {
			return false
		}
	}

	return true
}

// Size returns the dimension of the vector.
func (e Row) Size() int {
	return len(e)
}

// Dup returns a duplicate of this row.
func (e Row) Dup() Row {
	out := NewRow(e.Size())
	copy(out, e)

	return out
}

func (e Row) String() string {
	out := []rune{}
	out = append(out, []rune(fmt.Sprintf("%2.2x", []number.ByteFieldElem(e)))...)
	out = out[1 : len(out)-1]

	return string(out)
}
