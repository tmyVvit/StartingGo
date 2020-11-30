package leetcode

// 首先将nums1数组的数据后移，然后从前往后 双指针
// [1,2,3,0,0,0] --> [0,0,0,1,2,3]
//                          ^
// [4,5,6]
//  ^

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	for i := m - 1; i >= 0; i-- {
		nums1[n+i] = nums1[i]
	}
	id1, id2, k := n, 0, 0
	for id1 < m+n && id2 < n {
		if nums1[id1] <= nums2[id2] {
			nums1[k] = nums1[id1]
			id1++
			k++
		} else {
			nums1[k] = nums2[id2]
			id2++
			k++
		}
	}
	for id2 < n {
		nums1[k] = nums2[id2]
		id2++
		k++
	}
}

// 也可以直接从后往前：

// [1,2,3,0,0,0]
//      ^  -> ^
// [4,5,6]
//      ^
