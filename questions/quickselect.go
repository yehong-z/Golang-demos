package questions

import "math/rand"

func quickSelect(nums []int, k int) int {
	// 递归结束条件
	if len(nums) == 1 {
		return nums[0]
	}
	// 随机化 pivot 的位置
	pivotIndex := rand.Intn(len(nums))
	pivot := nums[pivotIndex]
	// 将数组划分为小于、等于和大于 pivot 的三部分
	var less, equal, greater []int
	for _, num := range nums {
		switch {
		case num < pivot:
			less = append(less, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			greater = append(greater, num)
		}
	}
	// 根据子数组长度的情况来递归查找第 k 小的元素
	switch {
	case k <= len(less):
		return quickSelect(less, k)
	case k > len(less)+len(equal):
		return quickSelect(greater, k-len(less)-len(equal))
	default:
		return equal[0]
	}
}

func quickselect2(arr []int, k int) int {
	l := make([]int, 0)
	r := make([]int, 0)
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] < arr[0] {
			l = append(l, arr[i])
		} else {
			r = append(r, arr[i])
		}
	}
	if len(l) == k-1 {
		return arr[0]
	} else if len(l) >= k {
		return quickselect2(l, k)
	} else {
		return quickselect2(r, k-len(l)-1)
	}
}
