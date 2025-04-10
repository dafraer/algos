func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	cnt := 0
	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			cnt++
		}
	}
	return cnt
}
