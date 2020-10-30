package leetcode

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}

type flag struct {
	node *TreeNode
	v    int
}

func myAppend(st []*flag, idx int, f *flag) []*flag {
	if len(st) == idx {
		st = append(st, f)
	} else {
		st[idx] = f
	}
	return st
}

// 使用迭代的方法，并且利用新的flag结构体
func preorderTraversalI(root *TreeNode) (res []int) {
	pointer := 0
	stack := []*flag{&flag{root, 0}}
	for pointer >= 0 {
		tmp := stack[pointer]
		if tmp.node == nil {
			pointer--
			continue
		}
		if tmp.v == 0 {
			// v==0表示刚入栈
			res = append(res, tmp.node.Val)
			tmp.v = 1
			pointer++
			stack = myAppend(stack, pointer, &flag{tmp.node.Left, 0})
		} else if tmp.v == 1 {
			// v==1 表示其左子节点已经遍历过了
			stack[pointer] = &flag{tmp.node.Right, 0}
		}
	}
	return res
}

// 也使用迭代的方法，利用一个栈
func PreorderTraversalII(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			//只要node不为空，就入栈
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		// node == nil时退出内循环，说明到了最左下角的节点
		// 找到栈顶的节点的右子节点，并更新栈（pop）
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
