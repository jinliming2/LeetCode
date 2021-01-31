package main

import "math/rand"

// Runtime: 56 ms, faster than 28.44% of Go online submissions for Insert Delete GetRandom O(1).
// Memory Usage: 10.2 MB, less than 8.26% of Go online submissions for Insert Delete GetRandom O(1).
// https://leetcode.com/submissions/detail/450173334/

// RandomizedSet :
type RandomizedSet struct {
	hash map[int]int
	arr  []int
}

// Constructor : Initialize your data structure here.
func Constructor() RandomizedSet {
	return RandomizedSet{
		hash: make(map[int]int),
		arr:  make([]int, 0),
	}
}

// Insert : Inserts a value to the set. Returns true if the set did not already contain the specified element.
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.hash[val]; ok {
		return false
	}
	rs.hash[val] = len(rs.arr)
	rs.arr = append(rs.arr, val)
	return true
}

// Remove : Removes a value from the set. Returns true if the set contained the specified element.
func (rs *RandomizedSet) Remove(val int) bool {
	if index, ok := rs.hash[val]; ok {
		last := len(rs.arr) - 1
		rs.arr[index], rs.arr[last] = rs.arr[last], rs.arr[index]
		rs.hash[rs.arr[index]] = index
		delete(rs.hash, val)
		rs.arr = rs.arr[:last]
		return true
	}
	return false
}

// GetRandom : Get a random element from the set.
func (rs *RandomizedSet) GetRandom() int {
	return rs.arr[rand.Intn(len(rs.arr))]
}
