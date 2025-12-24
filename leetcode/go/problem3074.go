import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sort.Ints(capacity)
	appleSum := 0
	for _, v := range apple {
		appleSum += v
	}
	i := len(capacity) - 1
	for i >= 0 && appleSum > 0 {
		appleSum -= capacity[i]
		i--
	}
	return len(capacity) - i - 1
}                         
