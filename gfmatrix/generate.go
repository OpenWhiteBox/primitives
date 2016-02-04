package gfmatrix

import (
	"io"

	"github.com/OpenWhiteBox/primitives/number"
)

// GenerateEmpty generates the n-by-m matrix with all entries set to 0.
func GenerateEmpty(n, m int) Matrix {
	out := make([]Row, n)

	for i := 0; i < n; i++ {
		out[i] = NewRow(m)
	}

	return out
}

// GenerateIdentity generates the n-by-n identity matrix.
func GenerateIdentity(n int) Matrix {
	out := GenerateEmpty(n, n)

	for i, _ := range out {
		out[i][i] = 0x01
	}

	return out
}

// GenerateRandomBinaryRow generates a random n-component row containing only 1s and 0s, using the random source reader.
func GenerateRandomBinaryRow(reader io.Reader, n int) Row {
	out := GenerateRandomRow(reader, n)

	for i, v := range out {
		out[i] = v & 1
	}

	return out
}

// GenerateRandomRow generates a random n-component row using the random source reader.
func GenerateRandomRow(reader io.Reader, n int) Row {
	out, temp := NewRow(n), make([]byte, n)
	reader.Read(temp)

	for i, v := range temp {
		out[i] = number.ByteFieldElem(v)
	}

	return Row(out)
}

// GenerateTrueRandom generates a random n-by-n matrix (not guaranteed to be invertible) using the random source reader
// (for example, crypto/rand.Reader).
func GenerateTrueRandom(reader io.Reader, n int) Matrix {
	m := make([]Row, n)

	for i := 0; i < n; i++ { // Generate random n x n matrix.
		m[i] = GenerateRandomRow(reader, n)
	}

	return m
}

// GenerateRandom generates a random invertible n-by-n matrix using the random source reader (for example,
// crypto/rand.Reader). Returns it and its inverse.
func GenerateRandom(reader io.Reader, n int) (Matrix, Matrix) {
	m := GenerateTrueRandom(reader, n)

	mInv, ok := m.Invert()
	if !ok { // Try again
		return GenerateRandom(reader, n) // Performance bottleneck.
	}

	return m, mInv // This one works!
}
