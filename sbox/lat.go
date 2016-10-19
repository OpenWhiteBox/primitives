package sbox

import (
	"fmt"

	"github.com/OpenWhiteBox/primitives/encoding"
)

var weight [4]uint64 = [4]uint64{
	0x6996966996696996, 0x9669699669969669,
	0x9669699669969669, 0x6996966996696996,
}

func dotProduct(a, b byte) byte {
	c := a & b
	return byte(weight[c/64]>>(c%64)) & 1
}

// LAT is a linear approximation table of a function.
type LAT [256][256]int

// NewLAT generates the LAT of a function.
func NewLAT(in encoding.Byte) *LAT {
	out := LAT{}

	for a := 0; a < 256; a++ {
		for b := 0; b < 256; b++ {
			for x := 0; x < 256; x++ {
				left := dotProduct(byte(a), in.Encode(byte(x)))
				right := dotProduct(byte(b), byte(x))

				if left == right {
					out[a][b]++
				}
			}

			out[a][b] -= 128
		}
	}

	return &out
}

func (lat *LAT) Linearity() int {
	out := 0

	for _, row := range lat[1:] {
		for _, col := range row {
			if col > out {
				out = col
			}
		}
	}

	return out
}

func (lat *LAT) Map() []int {
	max, min := 0, 0

	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			if lat[x][y] < min {
				min = lat[x][y]
			} else if lat[x][y] > max {
				max = lat[x][y]
			}
		}
	}

	out := make([]int, max-min+1)

	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			out[lat[x][y]-min]++
		}
	}

	return out
}

// Equals returns whether or not this LAT equals the given LAT.
func (lat *LAT) Equals(given *LAT) bool {
	for row := 0; row < 256; row++ {
		if lat[row] != given[row] {
			return false
		}
	}

	return true
}

// String serializes the LAT into a string.
func (lat *LAT) String() string {
	out := []rune{}

	for _, row := range lat {
		for _, col := range row {
			out = append(out, []rune(fmt.Sprintf("%v ", col))...)
		}
		out = append(out, '\n')
	}

	return string(out)
}
