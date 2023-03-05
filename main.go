package main

import (
	"fmt"
	"sort"
)

func main() {
	var horses int
	fmt.Scanf("%d", &horses)
	strengths := make([]int, 0, horses)
	for i := 0; i < horses; i++ {
		var strength int
		fmt.Scanf("%d", &strength)
		strengths = append(strengths, strength)
	}
	fmt.Println("str", strengths)
	sort.Ints(strengths)
	fmt.Println("str2", strengths)
	min := strengths[1] - strengths[0]
	for i := 2; i < horses; i++ {
		fmt.Printf("%d\n", strengths[i]-strengths[i-1])
		if dif := strengths[i] - strengths[i-1]; dif < min {
			min = dif
		}
	}
	fmt.Println("total min", min)
}
