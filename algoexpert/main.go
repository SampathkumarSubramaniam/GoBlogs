package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello")
	// TwoNumberSum([]int{1, 2, 3, 4}, 6)
	// SortedSquaredArray([]int{-1, 0, 1, 2})
	NonConstructibleChange([]int{5, 7, 1, 1, 2, 3, 22})
}

func NonConstructibleChange(coins []int) {
	i := 1
	sort.Ints(coins)
	for {
		if checkDirectlyInSlice(coins, i) {
			i++
			continue
		} else if checkStage1(coins, i) {
			i++
			continue
		} else if checkStage2(coins, i) {
			i++
			continue
		}
		break
	}
	fmt.Println("Not found", i)
}

func checkStage2(coins []int, exp int) bool {
	for i := 0; i < len(coins); i++ {
		for j := 0; j < len(coins); j++ {
			for k := 0; k < len(coins); k++ {
				if coins[i]+coins[j]+coins[k] == exp && i != j && i != k && j != k {
					return true
				}
			}
		}
	}
	return false

}

func checkStage1(coins []int, exp int) bool {
	for i := 0; i < len(coins); i++ {
		for j := 0; j < len(coins); j++ {
			if coins[i]+coins[j] == exp && i != j {
				return true
			}
		}
	}
	return false
}

func checkDirectlyInSlice(coins []int, exp int) bool {
	for i := 0; i < len(coins); i++ {
		if exp == coins[i] {
			return true
		}
	}
	return false
}

func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] != array[j] {
				if array[i]+array[j] == target {
					return []int{array[i], array[j]}
				}
			}
		}
	}
	return []int{}
}
func SortedSquaredArray(array []int) []int {
	// Write your code here.
	sort.Ints(array)
	var targetArray []int
	for i := 0; i < len(array); i++ {
		targetArray = append(targetArray, array[i]*array[i])
	}
	fmt.Println(targetArray)
	sort.Ints(targetArray)
	return targetArray
}

// coins:=[1,2,3,5] = 1,2,3,4,5,6,7,8,9,10,11
