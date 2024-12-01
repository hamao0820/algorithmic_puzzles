package main

import (
	"bytes"
	"context"
	"fmt"
	"math/bits"

	"github.com/goccy/go-graphviz"
)

func main() {
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

	ctx := context.Background()
	g, err := graphviz.New(ctx)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	graph, err := g.Graph(graphviz.WithDirectedType(graphviz.UnDirected))
	if err != nil {
		panic(err)
	}
	defer graph.Close()

	nodes := make(map[int]*graphviz.Node)
	for i := 0; i < 1<<4; i++ {
		if isIllegal[i] {
			continue
		}
		node, err := graph.CreateNodeByName(fmt.Sprintf("%04b", i))
		if err != nil {
			panic(err)
		}
		node.SetLabel(fmt.Sprintf("%04b", i))
		nodes[i] = node
	}

	// グラフを構築
	for i := 0; i < 1<<4; i++ {
		for j := i + 1; j < 1<<4; j++ {
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
				if _, err := graph.CreateEdgeByName(fmt.Sprintf("%d -> %d", i, j), nodes[i], nodes[j]); err != nil {
					panic(err)
				}
			}
		}
	}

	var buf bytes.Buffer
	if err := g.Render(ctx, graph, graphviz.XDOT, &buf); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
