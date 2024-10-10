package main

import "fmt"

func bubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	bubbleSort(s)
	fmt.Println(s)
}
