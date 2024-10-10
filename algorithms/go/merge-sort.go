package main

import "fmt"

func mergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	a := mergeSort(s[:len(s)/2])
	b := mergeSort(s[len(s)/2:])
	var i, j int
	ans := make([]int, 0, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			ans = append(ans, a[i])
			i++
		} else {
			ans = append(ans, b[j])
			j++
		}
	}
	if i >= len(a) {
		ans = append(ans, b[j:]...)
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	s = mergeSort(s)
	fmt.Println(s)

}
