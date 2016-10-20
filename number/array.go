package number

// ArrayRingElem is an element of Rijndael's ring, GF(2^8)[x]/(x^4 + 1).
//
// The additive identity is [0 0 0 0] and the multiplicative identity is [1 0 0 0].
type ArrayRingElem [4]ByteFieldElem

func NewArrayRingElem() ArrayRingElem {
	return ArrayRingElem{0, 0, 0, 0}
}

// Add returns e + f.
func (e ArrayRingElem) Add(f ArrayRingElem) ArrayRingElem {
	out := NewArrayRingElem()

	for i, _ := range out {
		out[i] = e[i].Add(f[i])
	}

	return out
}

// ScalarMul multiplies each component of e by a scalar from GF(2^8).
func (e ArrayRingElem) ScalarMul(g ByteFieldElem) ArrayRingElem {
	out := NewArrayRingElem()

	for i, e_i := range e {
		out[i] = e_i.Mul(g)
	}

	return out
}

// Mul returns e * f.
func (e ArrayRingElem) Mul(f ArrayRingElem) ArrayRingElem {
	out := NewArrayRingElem()

	for i, e_i := range e { // Foreach byte e_i in e:
		if !e_i.IsZero() { // with non-zero coefficient:
			temp := f.ScalarMul(e_i).shift(i) // Multiply f * e_i * x^i mod M(x):
			out = out.Add(temp)               // Add f * e_i * x^i to the output
		}
	}

	return out
}

func (e ArrayRingElem) shift(i int) ArrayRingElem {
	f := e.Dup()
	return ArrayRingElem{f[3-(i+3)%4], f[3-(i+2)%4], f[3-(i+1)%4], f[3-(i+0)%4]}
}

// Invert returns an element's multiplicative inverse, if it has one.
func (e ArrayRingElem) Invert() (ArrayRingElem, bool) {
	order := 256*256*256*256 - 1
	out, temp := ArrayRingElem{1, 0, 0, 0}, e.Dup()

	for i := uint(0); i < 32; i++ {
		if (order>>i)&1 == 1 {
			out = out.Mul(temp)
		}

		temp = temp.Mul(temp)
	}

	return out, out.Mul(e).IsOne()
}

// IsZero returns whether or not e is zero.
func (e ArrayRingElem) IsZero() bool { return e == ArrayRingElem{0, 0, 0, 0} }

// IsOne returns whether or not e is one.
func (e ArrayRingElem) IsOne() bool { return e == ArrayRingElem{1, 0, 0, 0} }

// Dup returns a duplicate of e.
func (e ArrayRingElem) Dup() ArrayRingElem {
	return ArrayRingElem{e[0].Dup(), e[1].Dup(), e[2].Dup(), e[3].Dup()}
}
