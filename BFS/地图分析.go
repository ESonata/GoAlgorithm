package BFS

/*	@ Main:多源BFS，记录所有陆地到其自身的距离为0，更新其到海洋的距离
	@ Href:./static/地图分析.png
*/

type node struct {
	x int
	y int
}

func maxDistance(grid [][]int) int {
	distance := make([][]int, 100+5)
	for i := 0; i < 100+5; i++ {
		distance[i] = make([]int, 100)
	}
	dx, dy := []int{1, -1, 0, 0}, []int{0, 0, 1, -1}
	n := len(grid)
	INF := int(1e6)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			distance[i][j] = INF
		}
	}
	queue := []node{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				distance[i][j] = 0
				queue = append(queue, node{i, j})
			}
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			nx := front.x + dx[i]
			ny := front.y + dy[i]
			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				if distance[nx][ny] > distance[front.x][front.y]+1 {
					distance[nx][ny] = distance[front.x][front.y] + 1
					queue = append(queue, node{nx, ny})
				}
			}
		}
	}

	ans := -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				ans = max(ans, distance[i][j])
			}
		}
	}
	if ans == INF {
		return -1
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
