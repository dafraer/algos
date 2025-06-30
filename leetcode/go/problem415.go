	func addStrings(num1 string, num2 string) string {
		a, _ := strconv.Atoi(num1)
		b, _ := strconv.Atoi(num2)
		return strconv.Itoa(a + b)
	}
