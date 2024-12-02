//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func main() {
	pairs := []int{5, 3, 2}
	n := len(pairs)

	gloves := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < pairs[i]; j++ {
			gloves = append(gloves, i+1)
			gloves = append(gloves, -(i + 1))
		}
	}
	sort.Slice(gloves, func(i, j int) bool {
		if gloves[i] < 0 && gloves[j] > 0 {
			return false
		}
		if gloves[i] > 0 && gloves[j] < 0 {
			return true
		}
		if gloves[i] < 0 && gloves[j] < 0 {
			return gloves[i] > gloves[j]
		}
		return gloves[i] < gloves[j]
	})

	sum := 0
	for _, p := range pairs {
		sum += 2 * p
	}
	pair := make(map[int]bool, n)
	for i := 1; i <= sum; i++ {
		if gloves[i-1] < 0 {
			pair[-gloves[i-1]] = true
		}
		if len(pair) == n {
			fmt.Println(i)
			return
		}
	}
}
