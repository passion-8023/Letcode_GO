package let_283

func moveZeroes(nums []int) {
	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != 0 {
			slow++
		} else {
			if nums[fast] != 0 {
				nums[slow], nums[fast] = nums[fast], nums[slow]
				slow++
			}
		}
	}
}

func moveZeroes1(nums []int) {
	if len(nums) == 0 {
		return
	}

	fast, slow := 0, 0
	for fast != len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	for slow != len(nums) {
		nums[slow] = 0
		slow++
	}

}
