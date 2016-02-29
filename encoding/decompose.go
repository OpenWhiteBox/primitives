package encoding

import (
	"crypto/rand"

	"github.com/OpenWhiteBox/primitives/matrix"
)

// EquivalentBytes returns true if two Byte encodings are identical and false if not.
func EquivalentBytes(a, b Byte) bool {
	for x := 0; x < 256; x++ {
		if a.Encode(byte(x)) != b.Encode(byte(x)) {
			return false
		}
	}

	return true
}

// DecomposeByteLinear decomposes an opaque Byte encoding into a ByteLinear encoding.
func DecomposeByteLinear(in Byte) (ByteLinear, bool) {
	m := matrix.Matrix{}
	for i := uint(0); i < 8; i++ {
		m = append(m, matrix.Row{in.Encode(byte(1 << i))})
	}

	forwards := m.Transpose()
	backwards, ok := forwards.Invert()

	return ByteLinear{
		Forwards:  forwards,
		Backwards: backwards,
	}, ok
}

// DecomposeByteAffine decomposes an opaque Byte encoding into a ByteAffine encoding.
func DecomposeByteAffine(in Byte) (ByteAffine, bool) {
	c := ByteAdditive(in.Encode(0))
	M, ok := DecomposeByteLinear(ComposedBytes{in, c})

	return ByteAffine{
		ByteLinear:   M,
		ByteAdditive: c,
	}, ok
}

// ProbablyEquivalentDoubles returns true if two Double encodings are probably equivalent and false if they're
// definitely not.
func ProbablyEquivalentDoubles(a, b Double) bool {
	for i := 0; i < 20; i++ {
		in := [2]byte{}
		rand.Read(in[:])

		x, y := a.Encode(in), b.Encode(in)
		if x != y {
			return false
		}
	}

	return true
}

// DecomposeDoubleLinear decomposes an opaque Double encoding into a DoubleLinear encoding.
func DecomposeDoubleLinear(in Double) (DoubleLinear, bool) {
	m := matrix.Matrix{}
	for i := 0; i < 2; i++ {
		for j := uint(0); j < 8; j++ {
			x := [2]byte{}
			x[i] = 1 << j
			x = in.Encode(x)

			m = append(m, matrix.Row(x[:]))
		}
	}

	forwards := m.Transpose()
	backwards, ok := forwards.Invert()

	return DoubleLinear{
		Forwards:  forwards,
		Backwards: backwards,
	}, ok
}

// DecomposeDoubleAffine decomposes an opaque Double encoding into a DoubleAffine encoding.
func DecomposeDoubleAffine(in Double) (DoubleAffine, bool) {
	c := DoubleAdditive(in.Encode([2]byte{}))
	M, ok := DecomposeDoubleLinear(ComposedDoubles{in, c})

	return DoubleAffine{
		DoubleLinear:   M,
		DoubleAdditive: c,
	}, ok
}

// ProbablyEquivalentWords returns true if two Word encodings are probably equivalent and false if they're definitely
// not.
func ProbablyEquivalentWords(a, b Word) bool {
	for i := 0; i < 20; i++ {
		in := [4]byte{}
		rand.Read(in[:])

		x, y := a.Encode(in), b.Encode(in)
		if x != y {
			return false
		}
	}

	return true
}

// DecomposeWordLinear decomposes an opaque Word encoding into a WordLinear encoding.
func DecomposeWordLinear(in Word) (WordLinear, bool) {
	m := matrix.Matrix{}
	for i := 0; i < 4; i++ {
		for j := uint(0); j < 8; j++ {
			x := [4]byte{}
			x[i] = 1 << j
			x = in.Encode(x)

			m = append(m, matrix.Row(x[:]))
		}
	}

	forwards := m.Transpose()
	backwards, ok := forwards.Invert()

	return WordLinear{
		Forwards:  forwards,
		Backwards: backwards,
	}, ok
}

// DecomposeWordAffine decomposes an opaque Word encoding into a WordAffine encoding.
func DecomposeWordAffine(in Word) (WordAffine, bool) {
	c := WordAdditive(in.Encode([4]byte{}))
	M, ok := DecomposeWordLinear(ComposedWords{in, c})

	return WordAffine{
		WordLinear:   M,
		WordAdditive: c,
	}, ok
}

// ProbablyEquivalentBlocks returns true if two Block encodings are probably equivalent and false if they're definitely
// not.
func ProbablyEquivalentBlocks(a, b Block) bool {
	for i := 0; i < 20; i++ {
		in := [16]byte{}
		rand.Read(in[:])

		x, y := a.Encode(in), b.Encode(in)
		if x != y {
			return false
		}
	}

	return true
}

// DecomposeBlockLinear decomposes an opaque Block encoding into a BlockLinear encoding.
func DecomposeBlockLinear(in Block) (BlockLinear, bool) {
	m := matrix.Matrix{}
	for i := 0; i < 16; i++ {
		for j := uint(0); j < 8; j++ {
			x := [16]byte{}
			x[i] = 1 << j
			x = in.Encode(x)

			m = append(m, matrix.Row(x[:]))
		}
	}

	forwards := m.Transpose()
	backwards, ok := forwards.Invert()

	return BlockLinear{
		Forwards:  forwards,
		Backwards: backwards,
	}, ok
}

// DecomposeBlockAffine decomposes an opaque Block encoding into a BlockAffine encoding.
func DecomposeBlockAffine(in Block) (BlockAffine, bool) {
	c := BlockAdditive(in.Encode([16]byte{}))
	M, ok := DecomposeBlockLinear(ComposedBlocks{in, c})

	return BlockAffine{
		BlockLinear:   M,
		BlockAdditive: c,
	}, ok
}

// DecomposeConcatenatedBlock decomposes an opaque concatenated Block encoding into an explicit one.
func DecomposeConcatenatedBlock(in Block) (out ConcatenatedBlock) {
	for pos := 0; pos < 16; pos++ {
		sbox := SBox{}

		for x := 0; x < 256; x++ {
			X := [16]byte{}
			X[pos] = byte(x)
			Y := in.Encode(X)

			sbox.EncKey[x] = Y[pos]
			sbox.DecKey[Y[pos]] = X[pos]
		}

		out[pos] = sbox
	}

	return
}
