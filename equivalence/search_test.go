package equivalence

import (
	"testing"

	"github.com/OpenWhiteBox/primitives/matrix"
)

func TestLearnConsistent(t *testing.T) {
	f := InvertEncoding{}
	A, B := matrix.NewDeductiveMatrix(8), matrix.NewDeductiveMatrix(8)

	for i := uint(0); i < 7; i++ {
		A.Assert(matrix.Row{byte(1 << i)}, matrix.Row{byte(1 << i)})
	}

	_, _, consistent := learn(f, f, A, B, 0, 0)
	if !consistent {
		t.Fatal("Learn said identity matrix was inconsistent.")
	}

	if !A.FullyDefined() {
		t.Fatal("Learn did not propagate knowledge into A.")
	} else if !B.FullyDefined() {
		t.Fatal("Learn did not propagate knowledge into B.")
	}

	if A.Matrix()[7][0] != 1<<7 {
		t.Fatal("Learn determined A incorrectly.")
	} else if B.Matrix()[7][0] != 1<<7 {
		t.Fatal("Learn determined B incorrectly.")
	}
}

func TestLearnInconsistent(t *testing.T) {
	f := InvertEncoding{}
	A, B := matrix.NewDeductiveMatrix(8), matrix.NewDeductiveMatrix(8)

	for i := uint(0); i < 6; i++ {
		A.Assert(matrix.Row{byte(1 << i)}, matrix.Row{byte(1 << i)})
	}
	A.Assert(matrix.Row{byte(1 << 7)}, matrix.Row{byte(1 << 6)})

	_, _, consistent := learn(f, f, A, B, 0, 0)

	if consistent {
		t.Fatal("Learn said inconsistent matrix was consistent.")
	}
}
