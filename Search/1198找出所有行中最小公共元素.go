package Search

func smallestCommonElement(mat [][]int) int {
	count := make([]int, 10005)
	m := len(mat)
	n := len(mat[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count[mat[i][j]]++
		}
	}

	for k := 1; k <= 10000; k++ {
		if count[k] == m {
			return k
		}
	}
	return -1
}
