package main

import "fmt"

/*

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

*/

func main() {
	numbs := []int{8, 0, 2, 11, 15, 7}

	for index, item := range numbs {
		for index2, item2 := range numbs[(index + 1):] {
			if (item + item2) == 9 {
				fmt.Println([]int{index, index2 + index + 1})
				break
			}
		}
	}

}
