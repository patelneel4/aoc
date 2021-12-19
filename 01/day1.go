package main

import (
    "fmt"
 	"io/ioutil"
 	"strings"
	"strconv"
)



func readFile(fname string) (nums []int, err error) {
	//Credit- https://stackoverflow.com/a/9863218/6642224
    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }

    lines := strings.Split(string(b), "\n")
    // Assign cap to avoid resize on every append.
    nums = make([]int, 0, len(lines))

    for _, l := range lines {
        // Empty line occurs at the end of the file when we use Split.
        if len(l) == 0 { continue }
        // Atoi better suits the job when we know exactly what we're dealing
        // with. Scanf is the more general option.
        n, err := strconv.Atoi(l)
        if err != nil { return nil, err }
        nums = append(nums, n)
    }

    return nums, nil
}
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
	nums, err := readFile("input.txt")
    if err != nil { panic(err) }
	fmt.Println("Part 1: " , Part1(nums))
	fmt.Println("Part 2: " , Part2(nums))
	


}
