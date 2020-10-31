package leetcode

import "math/rand"

type RandomizedSet struct {
	table map[int]int
	list  []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), []int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.table[val]; ok {
		return false
	}
	this.table[val] = len(this.list)
	this.list = append(this.list, val)
	return true

}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.table[val]; !ok {
		return false
	}
	tmp := this.table[val]
	last := this.list[len(this.list)-1]
	this.table[last] = tmp
	this.list[tmp] = last
	this.list = this.list[:len(this.list)-1]
	this.table[val] = -1
	return true

}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
