package main

func maxProduct(words []string) int {
	m := map[int]int{}

	for k, v := range words {
		for i := 0; i < len(v); i++ {
			index := v[i] - 'a'
			m[k] = m[k] | (1 << index)
		}
	}
	maxnum := 0
	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(m); j++ {
			if m[i]&m[j] == 0 && len(words[i])*len(words[j]) > maxnum {
				maxnum = len(words[i]) * len(words[j])
			}
		}
	}

	return maxnum

}

func main() {

}
