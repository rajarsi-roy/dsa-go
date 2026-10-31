/*

Problem Url: https://leetcode.com/problems/two-sum/

Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.

*/

package main

import "fmt"

func twoSumPresence(nums []int, target int) (targetInNums bool) {
	uniqueElements := make(map[int]struct{})
	for _, element := range nums {
		if _, isPresentInUniqueElements := uniqueElements[target - element]; isPresentInUniqueElements {
			targetInNums = true
			break
		}
		uniqueElements[element] = struct{}{}
	}
	return targetInNums
}

func twoSum(nums []int, target int) (indexPair []int) {
	elementToIndexMap := make(map[int]int)
	for index, element := range nums {
		indexOfComplement, isComplementPresent := elementToIndexMap[target - element]
		if isComplementPresent {
			indexPair = []int{index, indexOfComplement}
			break
		}
		elementToIndexMap[element] = index
	}
	return indexPair
}

func main() {
	fmt.Println(twoSumPresence([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
}
