package main

import (
	utils "aoc/utils"
	"fmt"

)

func Part1(nums []int) int{
	var iVal = nums[0]
	counter:=0
	for i := 0; i < len(nums); i++ {
		if nums[i] > iVal{ 
			counter++
			iVal = nums[i]
		} else {
			iVal = nums[i]
		}
		
	}

	return counter
}
func Part2(nums []int) int{
	var iVal = nums[0]
	counter:=0
	for i := 2; i < len(nums); i++ {
		var i1 = nums[i]
		var iTtl = i1
		if (i+2) < len(nums){
			fmt.Println(i)
			var i2 = nums[i + 1]
			var i3 = nums[i + 2]
			iTtl = i1+i2+i3
		}
		
		if iTtl > iVal{ 
			counter++
		}
		iVal = iTtl
		
	}

	return counter
}

func main() {
	nums, err := utils.ReadIntArrayFromFile("input.txt")
    if err != nil { panic(err) }
	fmt.Println("Part 1: " , Part1(nums))
	fmt.Println("Part 2: " , Part2(nums))
	


}
