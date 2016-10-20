// Package equivalence implements the linear equivalence algorithm. TODO: The affine equivalence algorithm.
package equivalence

import (
	"github.com/OpenWhiteBox/primitives/encoding"
	"github.com/OpenWhiteBox/primitives/matrix"
)

// Additive is an additive equivalence--a pair of bytes, (A, B) such that f(x + a) = f(x) + b for all x.
type Additive struct {
	A, B encoding.ByteAdditive
}

// Linear is a linear equivalence--a pair of matrices, (A, B) such that f(A(x)) = B(g(x)) for all x.
type Linear struct {
	A, B encoding.ByteLinear
}

// Affine is an affine equivalence--a part of affine transformations such that f(A(x)) = B(g(x)) for all x.
type Affine struct {
	A, B encoding.ByteAffine
}

// FindAdditive finds additive equivalences between f and g. cap is the maximum number of equivalences to return.
func FindAdditive(f, g encoding.Byte, cap int) (out []Additive) {
	for a := 0; a < 256 && len(out) < cap; a++ {
		b := f.Encode(byte(a)) ^ g.Encode(0)
		ok := true

		for x := 1; x < 256 && ok; x++ {
			ok = ok && b == f.Encode(byte(x)^byte(a))^g.Encode(byte(x))
		}

		if ok {
			out = append(out, Additive{
				A: encoding.ByteAdditive(a),
				B: encoding.ByteAdditive(b),
			})
		}
	}

	return out
}

// FindLinear finds linear equivalences between f and g. cap is the maximum number of equivalences to return.
func FindLinear(f, g encoding.Byte, cap int) []Linear {
	return search(f, g, matrix.NewDeductiveMatrix(8), matrix.NewDeductiveMatrix(8), 0, 0, cap)
}

// FindAffine finds affine equivalences between f and g. cap is the maximum number of equivalences to return. (Not
// Implemented.)
func FindAffine(f, g encoding.Byte, cap int) []Affine {
	return nil
}
