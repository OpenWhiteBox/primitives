// Package sbox implements several general purpose tools for investigating S-boxes.
package sbox

type SBoxEncoding [16][16]byte

func (sbe SBoxEncoding) Encode(in byte) byte {
	return sbe[in>>4][in&0x0f]
}

func (sbe SBoxEncoding) Decode(in byte) byte {
	for upper := 0; upper < 16; upper++ {
		for lower := 0; lower < 16; lower++ {
			if sbe[upper][lower] == in {
				return byte(upper)<<4 | byte(lower)
			}
		}
	}
	return 0
}
