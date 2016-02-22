package gfmatrix

// RightStretch returns the matrix of right multiplication by the given matrix.
func (e Matrix) RightStretch() Matrix {
	n, m := e.Size()
	nm := n * m

	out := GenerateEmpty(nm, nm)

	for i := 0; i < nm; i++ {
		p, q := i/n, i%n

		for j := 0; j < m; j++ {
			out[i][j*m+q] = e[p][j]
		}
	}

	return out
}

// LeftStretch returns the matrix of left matrix multiplication by the given matrix.
func (e Matrix) LeftStretch() Matrix {
	n, m := e.Size()
	nm := n * m

	out := GenerateEmpty(nm, nm)

	for i := 0; i < nm; i++ {
		p, q := i/n, i%n

		for j := 0; j < m; j++ {
			out[i][j+m*p] = e[j][q]
		}
	}

	return out
}
