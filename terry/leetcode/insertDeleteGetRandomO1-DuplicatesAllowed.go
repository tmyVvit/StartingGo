package leetcode

import "math/rand"

// RandomizedCollection is a struct that can insert, remove and get rancom number in it
// it can store duplicate values
type RandomizedCollection struct {
	table map[int][]int // table中存放所有数据在list中的index
	list  []int         // list中存放所有集合中的数据，包括重复的数据
}

// NewRandomizedCollection initial an empty collection
func NewRandomizedCollection() *RandomizedCollection {
	return &RandomizedCollection{
		table: make(map[int][]int),
		list:  make([]int, 0)}
}

// Insert value to collection
// if collection contains val return false, else return true
// but insert action always success
func (rColl *RandomizedCollection) Insert(val int) (ret bool) {
	if _, ok := rColl.table[val]; !ok {
		rColl.table[val] = make([]int, 0)
	}
	rColl.list = append(rColl.list, val)
	rColl.table[val] = append(rColl.table[val], len(rColl.list)-1)
	return len(rColl.table[val]) == 1
}

// Remove a value from the collection. Returns true if the collection contained the specified element.
func (rColl *RandomizedCollection) Remove(val int) bool {
	tval, ok := rColl.table[val]
	if !ok || len(tval) == 0 {
		return false
	}
	// 首先更新table中val的index
	rmIdx := tval[len(tval)-1]
	rColl.table[val] = tval[:len(tval)-1]
	// 将list中最后的元素和要remove的元素交换位置，然后删除最后的元素
	rColl.list[rmIdx] = rColl.list[len(rColl.list)-1]
	// 此时rColl.list[len(rColl.list)-1]将会移到list的rmIdx下标处，所以还要更新table中的index
	rColl.table[rColl.list[len(rColl.list)-1]] = updateTable(rColl.table[rColl.list[len(rColl.list)-1]], len(rColl.list)-1, rmIdx)
	// 将list最后的元素移除
	rColl.list = rColl.list[:len(rColl.list)-1]
	return true
}

// update table中的下标时需要遍历整个下标的list，时间复杂度时O(n)
// 可以考虑使用 RandomizedSet 来存放相同数字的下标
func updateTable(itable []int, old int, new int) []int {
	for id, val := range itable {
		if val == old {
			itable[id] = new
			return itable
		}
	}
	return itable
}

// GetRandom Get a random element from the collection.
func (rColl *RandomizedCollection) GetRandom() int {
	return rColl.list[rand.Intn(len(rColl.list))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
