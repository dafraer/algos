package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		if len(s) <= 10 {
			fmt.Println(s)
		}
		var ans strings.Builder
		ans.WriteByte(s[0])
		ans.WriteString(strconv.Itoa())
		ans.WriteByte(s[len(s)-1])
		ans
	}
}
