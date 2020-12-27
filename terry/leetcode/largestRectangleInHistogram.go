package leetcode

// 以高度循环遍历，找到此高度下的最大面积 时间复杂度O(n^2)
func largestRectangleArea(heights []int) int {
	length := len(heights)
	area := 0
	for i := 0; i < length; i++ {
		left, right := i, i
		for left-1 >= 0 && heights[left-1] >= heights[i] {
			left--
		}
		for right+1 < length && heights[right+1] >= heights[i] {
			right++
		}
		area = max(area, heights[i]*(right-left+1))
	}
	return area
}

/**
其实上面的方法就是找到当前坐标的两边第一个小于当前高度的坐标，分别设为左边界 右边界
根据单调栈 从左到右 从右到左 遍历两次，分别寻找左边界 右边界， 最后再遍历一遍计算面积
时间复杂度O(n)
*/
func largestRectangleAreaI(heights []int) int {
	length := len(heights)
	left := make([]int, length)
	right := make([]int, length)
	stack := make([]int, length)
	idx := 0
	for i := 0; i < length; i++ {
		// 如果栈顶的元素大于单调当前元素值，则出栈
		for idx > 0 && heights[stack[idx-1]] >= heights[i] {
			idx--
		}
		// 如果左边没有比当前元素小的，那么另当前元素的左边界为 -1
		if idx == 0 {
			left[i] = -1
		} else {
			left[i] = stack[idx-1]
		}
		stack[idx] = i
		idx++
	}
	idx = 0
	for i := length - 1; i >= 0; i-- {
		for idx > 0 && heights[stack[idx-1]] >= heights[i] {
			idx--
		}
		if idx == 0 {
			right[i] = length
		} else {
			right[i] = stack[idx-1]
		}
		stack[idx] = i
		idx++
	}
	area := 0
	for i := 0; i < length; i++ {
		area = max(area, heights[i]*(right[i]-left[i]-1))
	}
	return area
}
