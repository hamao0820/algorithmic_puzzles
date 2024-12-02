//go:build ignore

package main

import "fmt"

func main() {
	path := make([]int, 0)

	// 許されない組み合わせを列挙
	isIllegal := make(map[int]bool)
	for i := 0; i < 1<<4; i++ {
		// 狼と羊が同じ
		if (i>>2&1) == (i>>1&1) && (i>>3) != (i>>2&1) {
			isIllegal[i] = true
		}
		// 羊とキャベツが同じ
		if (i>>1&1) == (i&1) && (i>>3) != (i>>1&1) {
			isIllegal[i] = true
		}
	}

	visited := make([]bool, 1<<4)
	var dfs func(int) bool
	dfs = func(v int) bool {
		if v == (1<<4)-1 {
			path = append(path, v)
			return true
		}
		visited[v] = true
		path = append(path, v)
		for i := 0; i < 4; i++ {
			next := v ^ (1 << 3)
			if i < 3 {
				next ^= 1 << i
			}
			if isIllegal[next] {
				continue
			}
			if visited[next] {
				continue
			}
			if dfs(next) {
				return true
			}
		}
		path = path[:len(path)-1]
		visited[v] = false
		return false
	}

	dfs(0)
	for _, v := range path {
		fmt.Printf("%04b\n", v)
	}
}
