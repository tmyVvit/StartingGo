package leetcode

// GetAllElements 合并两个二叉搜索树并排序
// 首先将两个tree转化成两个数组
// 然后数组合并
func GetAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := treeToSlice(root1)
	arr2 := treeToSlice(root2)
	return mergeSorted(arr1, arr2)
}

func treeToSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := treeToSlice(root.Left)
	left = append(left, root.Val)
	return append(left, treeToSlice(root.Right)...)
}

func mergeSorted(arr1 []int, arr2 []int) (res []int) {
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		res = append(res, arr1[i])
		i++
	}
	for j < len(arr2) {
		res = append(res, arr2[j])
		j++
	}
	return res
}

// GetAllElementsI 使用goroutine实现
// 同时对两棵树进行中序排列，并将结果传递到两个channel中
// 读取两个channel的数据
func GetAllElementsI(root1 *TreeNode, root2 *TreeNode) []int {
	ch1, ch2 := make(chan int), make(chan int)
	out := make(chan int)
	res := []int{}
	if root1 != nil {
		go treeValToChannel(root1, root1.Val, ch1)
	} else {
		close(ch1)
	}
	if root2 != nil {
		go treeValToChannel(root2, root2.Val, ch2)
	} else {
		close(ch2)
	}

	go func() {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		for ok1 || ok2 {
			if ok1 && (v1 < v2 || !ok2) {
				out <- v1
				v1, ok1 = <-ch1
			} else {
				out <- v2
				v2, ok2 = <-ch2
			}
		}
		close(out)
	}()
	for val := range out {
		res = append(res, val)
	}
	return res
}

func treeValToChannel(root *TreeNode, rootVal int, ch chan int) {
	if root == nil {
		return
	}
	treeValToChannel(root.Left, rootVal, ch)
	ch <- root.Val
	treeValToChannel(root.Right, rootVal, ch)
	if rootVal == root.Val {
		close(ch)
	}
}
