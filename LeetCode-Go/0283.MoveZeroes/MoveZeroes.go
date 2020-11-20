package leetcode

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	for i := 0; i < len(nums); i++ {
		j := i + 1
		if j > len(nums)-1 {
			j = len(nums) - 1
		}
		for nums[i] == 0 {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			} else {
				j++
			}
			if j == len(nums)-1 && nums[i] == 0 && nums[j] == 0 {
				break
			}
		}
	}
}
