package let_88

import (
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	//方法一：最low的方法(最优解)
	//for i := 0; i < n; i++ {
	//	nums1[m+i] = nums2[i]
	//}
	//for i := 0; i < m+n; i++ {
	//	for j := i+1; j < m+n; j++ {
	//		if nums1[i] > nums1[j] {
	//			nums1[i], nums1[j] = nums1[j], nums1[i]
	//		}
	//	}
	//}
	//return nums1

	//方法二：copy sort
	//copy(nums1[m:], nums2)
	//sort.Ints(nums1)

	//方法三：双指针
	//result := make([]int, 0,  m+n)
	//p1, p2 := 0, 0
	//for {
	//	if p1 == m {
	//		result = append(result, nums2[p2:]...)
	//		break
	//	}
	//	if p2 == n {
	//		result = append(result, nums1[p1:]...)
	//		break
	//	}
	//	if nums1[p1] > nums2[p2] {
	//		result = append(result, nums2[p2])
	//		p2++
	//	} else {
	//		result = append(result, nums1[p1])
	//		p1++
	//	}
	//}
	//copy(nums1, result)
	//return nums1

	//方法四：逆向双指针
	for p1, p2, tail :=  m-1, n-1, m+n-1; p1 >=0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
	return nums1
}

func TestMerge(t *testing.T)  {
	nums1 := []int{4,5,6,0,0,0}
	m := 3
	nums2 := []int{1,2,3}
	n := 3

	result := merge(nums1, m, nums2, n)
	t.Log(result)
}
