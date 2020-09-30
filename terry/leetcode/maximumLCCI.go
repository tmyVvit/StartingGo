package leetcode

// func main() {
// 	fmt.Println(maximum(2000, 11111))
// }

// 不使用 if-else 或者任何比较运算符
func maximum(a int, b int) int {
	x := int64(a) - int64(b)
	k := (x >> 63) & 1
	return a*(1-int(k)) + b*int(k)
}
