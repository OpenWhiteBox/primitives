package sbox

import (
	"fmt"

	"github.com/OpenWhiteBox/primitives/encoding"
)

// DDT is a difference distribution table of a function.
type DDT [256][256]int

// NewDDT generates the DDT of a function.
func NewDDT(in encoding.Byte) *DDT {
	out := DDT{}
	out[0][0] = 256

	for a := 1; a < 256; a++ {
		for x := 0; x < 256; x++ {
			b := in.Encode(byte(x)^byte(a)) ^ in.Encode(byte(x))
			out[a][b]++
		}
	}

	return &out
}

// Uniform returns the differential uniformity of the DDT--the largest entry that isn't at position (0, 0).
func (ddt *DDT) Uniform() int {
	out := 0

	for _, row := range ddt[1:] {
		for _, col := range row {
			if col > out {
				out = col
			}
		}
	}

	return out
}

// Equals returns whether or not this DDT equals the given DDT.
func (ddt *DDT) Equals(given *DDT) bool {
	for row := 0; row < 256; row++ {
		if ddt[row] != given[row] {
			return false
		}
	}

	return true
}

// String serializes the DDT to a string.
func (ddt *DDT) String() string {
	out := []rune{}

	for _, row := range ddt {
		for _, col := range row {
			out = append(out, []rune(fmt.Sprintf("%v ", col))...)
		}
		out = append(out, '\n')
	}

	return string(out)
}
