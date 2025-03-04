func finalValueAfterOperations(operations []string) int {
	var x int8
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "--X":
			x--
		case "X--":
			x--
		case "++X":
			x++
		case "X++":
			x++
		}
	}
	return int(x)
}
