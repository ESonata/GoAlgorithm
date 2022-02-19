package main

func longestLine(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	a := make([][]int, m)
	b := make([][]int, m)
	c := make([][]int, m)
	d := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
		b[i] = make([]int, n)
		c[i] = make([]int, n)
		d[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				a[i][j] = 0
				b[i][j] = 0
				c[i][j] = 0
				d[i][j] = 0
			} else {
				if j > 0 {
					a[i][j] = a[i][j-1] + 1
				} else {
					a[i][j] = 1
				}

				if i > 0 {
					b[i][j] = b[i-1][j] + 1
				} else {
					b[i][j] = 1
				}

				if i > 0 && j > 0 {
					c[i][j] = c[i-1][j-1] + 1
				} else {
					c[i][j] = 1
				}

				if i > 0 && j < n-1 {
					d[i][j] = d[i-1][j+1] + 1
				} else {
					d[i][j] = 1
				}
				ans = max(ans, a[i][j])
				ans = max(ans, b[i][j])
				ans = max(ans, c[i][j])
				ans = max(ans, d[i][j])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
