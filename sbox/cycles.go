package sbox

import (
	"github.com/OpenWhiteBox/primitives/encoding"
)

// ByteCycles contains the cycles of a Byte encoding.
type ByteCycles [][]byte

// NewByteCycles takes a Byte encoding as input and returns the corresponding ByteCycles object.
func NewByteCycles(in encoding.Byte) ByteCycles {
	out := ByteCycles{}

	unused := make(map[byte]bool)
	for i := 0; i < 256; i++ {
		unused[byte(i)] = true
	}

	for next, _ := range unused {
		cycle := []byte{}

		state := next
		for {
			cycle = append(cycle, state)

			delete(unused, state)
			state = in.Encode(state)

			if state == next {
				break
			}
		}

		out = append(out, cycle)
	}

	return out
}
