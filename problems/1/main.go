package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// グラフを隣接リストで表現
	graph := make([][]int, 1<<4)

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

	// グラフを構築
	for i := 0; i < 1<<4; i++ {
		for j := 0; j < 1<<4; j++ {
			if isIllegal[i] || isIllegal[j] {
				continue
			}

			// 4ビット目が等しい場合は隣接していない
			if (i >> 3) == (j >> 3) {
				continue
			}

			// 3ビット以上異なる場合は隣接していない
			if bits.OnesCount(uint(i^j)) > 2 {
				continue
			}

			if i>>3 == (i >> bits.TrailingZeros(uint(i^j)) & 1) {
				graph[i] = append(graph[i], j)
			}
		}
	}

	// 0000から1111までの最短経路を求める
	dist := make([]int, 1<<4)
	pre := make([]int, 1<<4)
	for i := range dist {
		dist[i] = -1
		pre[i] = -1
	}
	dist[0] = 0
	queue := []int{0}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, u := range graph[v] {
			if dist[u] == -1 {
				dist[u] = dist[v] + 1
				pre[u] = v
				queue = append(queue, u)
			}
		}
	}

	// 1111から0000までの最短経路を復元
	path := []int{}
	for v := 1<<4 - 1; v != 0; v = pre[v] {
		path = append(path, v)
	}
	path = append(path, 0)

	// 結果を出力
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf("%04b\n", path[i])
	}
}
