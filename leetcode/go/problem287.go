func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	for true {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			newSlow := 0
			for newSlow != slow {
				slow = nums[slow]
				newSlow = nums[newSlow]
			}
			break
		}
	}
	return slow
}
