package let_33

func Search(nums []int, target int) int {
	n := len(nums)
	maxi := 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			maxi = i - 1
			break;
		}
	}
	nums1 := nums[:maxi+1]
	nums2 := nums[maxi+1:]
	low := 0
	high := len(nums1) - 1
	for low <= high {
		var mid int = low + ((high - low) >> 1)
		if nums1[mid] > target {
			high = mid - 1
		} else if nums1[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}

	low = 0
	high = len(nums2) - 1
	for low <= high {
		var mid int = low + ((high - low) >> 1)
		if nums2[mid] > target {
			high = mid - 1
		} else if nums2[mid] < target {
			low = mid + 1
		} else {
			return mid + len(nums1)
		}
	}
	return -1
}



