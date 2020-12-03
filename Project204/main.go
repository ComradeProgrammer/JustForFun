package main

import (
	"./leetcode"
	"fmt"
)

func main() {
	var a []int = []int{269826447, 974181916, 225871443, 189215924, 664652743, 592895362, 754562271, 335067223, 996014894, 510353008, 48640772, 228945137}
	var b = 3
	var r = leetcode.MaxDistance(a, b)
	fmt.Println(r)

}
