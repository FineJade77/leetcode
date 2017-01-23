package main

import "fmt"

type CustomMap map[int]int

func (c CustomMap) isContains(key int) bool {
	if _, ok := c[key]; ok {
		return true
	}
	return false
}

func twoSum(nums []int, target int) []int {
	numMap := CustomMap(make(map[int]int))
	for ix, val := range nums {
		if numMap.isContains(val) {
			return []int{numMap[val], ix}
		} else {
			numMap[target-val] = ix
		}
	}
	return []int{0, 0}
}

func main() {
	nums := []int{0, 2, 6, 3}
	ret := twoSum(nums, 5)
	fmt.Println(ret)
}
