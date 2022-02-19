package main

func deleteTreeNodes(nodes int, parent []int, value []int) int {
	edge := make([][]int, nodes)

	for i := 0; i < len(parent); i++ {
		if parent[i] != -1 {
			edge[parent[i]] = append(edge[parent[i]], i)
		}
	}
	cnt := make([]int, nodes)
	for i := 0; i < len(cnt); i++ {
		cnt[i] = 1
	}
	var dfs func(now int)

	dfs = func(now int) {
		for i := 0; i < len(edge[now]); i++ {
			dfs(edge[now][i])
			value[now] += value[edge[now][i]]
			cnt[now] += cnt[edge[now][i]]
		}
		if value[now] == 0 {
			cnt[now] = 0
		}
	}

	dfs(0)

	return cnt[0]
}
