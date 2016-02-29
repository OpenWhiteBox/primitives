package equivalence

import (
	"testing"

	"crypto/rand"

	"github.com/OpenWhiteBox/primitives/encoding"
	"github.com/OpenWhiteBox/primitives/matrix"
	"github.com/OpenWhiteBox/primitives/number"
)

type InvertEncoding struct{}

func (ie InvertEncoding) Encode(in byte) byte { return byte(number.ByteFieldElem(in).Invert()) }
func (ie InvertEncoding) Decode(in byte) byte { return ie.Encode(in) }

func TestFindAdditive(t *testing.T) {
	L := encoding.NewByteLinear(matrix.GenerateRandom(rand.Reader, 8))
	eqs := FindAdditive(L, L, 20)

	for _, eq := range eqs {
		if L.Encode(byte(eq.A)) != byte(eq.B) {
			t.Fatal("FindAdditive found an incorrect equivalence!")
		}
	}
}

func TestFindLinear(t *testing.T) {
	cap := 2041
	if testing.Short() {
		cap = 20
	}

	f := InvertEncoding{}
	eqs := FindLinear(f, f, cap)

	if !(testing.Short() && len(eqs) == 20 || !testing.Short() && len(eqs) == 2040) {
		t.Fatalf("FindLinear found the wrong number of equivalences! Wanted %v, got %v.", 2040, len(eqs))
	}

	for _, eq := range eqs {
		fA := encoding.ComposedBytes{eq.A, f}
		Bf := encoding.ComposedBytes{f, eq.B}

		if !encoding.EquivalentBytes(fA, Bf) {
			t.Fatalf("FindLinear found an incorrect equivalence.")
		}
	}
}
