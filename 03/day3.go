package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)


func Part1(lines []string) {
	gamma, epsilon := "", ""
	for i:=0; i <len(lines[0]); i++{
		count0, count1 := 0,0
		for _, line := range lines{
			if line[i] == '0'{
				count0++
			} else {
				count1++
			}
		}
		if(count0 > count1){
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gammaInt, _ := strconv.ParseInt(gamma,2,64)
	epsilonInt, _ := strconv.ParseInt(epsilon,2,64)
	fmt.Println(int(gammaInt) * int(epsilonInt))
}
// func Part2(sc *bufio.Scanner) int{
	
// }

func main() {
	lines, _:= utils.ReadStringArrayFromFile("input.txt")
	Part1(lines)
}
