package number

import (
	"testing"
)

func TestByteFieldElemMul(t *testing.T) {
	x, y := ByteFieldElem(0x57), ByteFieldElem(0x83)

	if x.Mul(y) != 0xc1 || y.Mul(x) != 0xc1 {
		t.Fatalf("0x57 * 0x83 != 0xc1")
	}
}

func TestByteFieldElemInvert(t *testing.T) {
	for w := 1; w < 256; w++ {
		x := ByteFieldElem(w)
		y := x.Invert()

		if x.Mul(y) != 0x01 || y.Mul(x) != 0x01 {
			t.Fatalf("Multiplication of byte element by inverse did not equal one!")
		}
	}
}

func TestFactor(t *testing.T) {
	for w := 1; w < 256; w++ {
		x := ByteFieldElem(w)

		y := ByteFieldElem(w).Mul(x).Mul(x).Mul(x)
		if y == 255 {
			t.Log(x)
		}
	}
}

func TestArrayRingElemMultiplicationArbitrary(t *testing.T) {
	x := ArrayRingElem{0x02, 0x01, 0x01, 0x03}
	y := ArrayRingElem{0x0e, 0x09, 0x0d, 0x0b}

	if !x.Mul(y).IsOne() || !y.Mul(x).IsOne() {
		t.Fatalf("Multiplication is wrong, element * inverse != 1")
	}
}

func TestArrayRingElemMultiplicationOne(t *testing.T) {
	x := ArrayRingElem{0x02, 0x01, 0x01, 0x03}
	y := ArrayRingElem{0x01, 0x00, 0x00, 0x00}

	xy, yx := x.Mul(y), y.Mul(x)

	for i := 0; i < 4; i++ {
		if xy[i] != x[i] || yx[i] != x[i] {
			t.Fatalf("Multiplication is wrong, element * 1 != element")
		}
	}
}

func TestArrayRingElemMultiplicationZero(t *testing.T) {
	x := ArrayRingElem{0x02, 0x01, 0x01, 0x03}
	y := ArrayRingElem{0x00, 0x00, 0x00, 0x00}

	if !x.Mul(y).IsZero() || !y.Mul(x).IsZero() {
		t.Fatalf("Multiplication is wrong, element * 0 != 0")
	}
}

func TestArrayRingElemMultiplicationInvert(t *testing.T) {
	x := ArrayRingElem{0x02, 0x01, 0x01, 0x03}
	y := ArrayRingElem{0x00, 0x00, 0x00, 0x00}

	if _, ok := x.Invert(); !ok {
		t.Fatal("Invert is wrong, failed to find inverse of unit.")
	}

	if _, ok := y.Invert(); ok {
		t.Fatal("Invert is wrong, found inverse of non-unit.")
	}
}
