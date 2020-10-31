package leetcode

import "math/rand"

// RandomizedSet it a struct that can insert, remove and get rancom number in it
// with O(1) in time
type RandomizedSet struct {
	table map[int]int
	list  []int
}

// NewRandomSet return a new empty RandomizedSet
func NewRandomSet() *RandomizedSet {
	return &RandomizedSet{make(map[int]int), []int{}}
}

// Insert a val into set
// return true if insert success, false if failed
func (rset *RandomizedSet) Insert(val int) bool {
	if _, ok := rset.table[val]; ok {
		return false
	}
	rset.table[val] = len(rset.list)
	rset.list = append(rset.list, val)
	return true

}

// Remove a val in the set
// true if remove success and false if failed
func (rset *RandomizedSet) Remove(val int) bool {
	if _, ok := rset.table[val]; !ok {
		return false
	}
	tmp := rset.table[val]
	last := rset.list[len(rset.list)-1]
	rset.table[last] = tmp
	rset.list[tmp] = last
	rset.list = rset.list[:len(rset.list)-1]
	rset.table[val] = -1
	return true

}

// GetRandom can get a random number in set
func (rset *RandomizedSet) GetRandom() int {
	return rset.list[rand.Intn(len(rset.list))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
