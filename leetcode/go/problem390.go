func lastRemaining(n int) int {
	head, step := 1, 1
	left := true

	for n > 1 {
		if left || n%2 == 1 {
			head += step
		}

		n /= 2
		step *= 2
		left = !left
	}

	return head
}