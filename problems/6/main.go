package main

import "fmt"

func main() {
	n := 1000

	switch n % 8 {
	case 1:
		fmt.Println("親指")
	case 2, 0:
		fmt.Println("人差し指")
	case 3, 7:
		fmt.Println("中指")
	case 4, 6:
		fmt.Println("薬指")
	}
}
