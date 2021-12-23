package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)
func mostCommonAndLeastBit(lines [] string, i int) (string, string){
	count0, count1 := 0,0
	for _, line := range lines{
		if line[i] == '0'{
			count0++
		} else {
			count1++
		}
	}
	if(count0 > count1){
		return "0", "1"
	} else {
		return "1", "0"
	}

}

func part1(lines []string) {
	gamma, epsilon := "",""
	for i:=0; i <len(lines[0]); i++{
		g, e :=  mostCommonAndLeastBit(lines, i)
		gamma += g
		epsilon += e
	}
	gammaInt, _ := strconv.ParseInt(gamma,2,64)
	epsilonInt, _ := strconv.ParseInt(epsilon,2,64)
	fmt.Println(int(gammaInt) * int(epsilonInt))
}
func filteredLines(lines []string, i int, isCommon bool) string{
	if len(lines) == 1{
		return lines[0]
	}
	most, least := mostCommonAndLeastBit(lines, i)
	comparer := ""
	if isCommon {
		comparer = most
	} else {
		comparer = least
	}
	filteredList := make([]string, 0)
	for _, line := range lines {
		if string(line[i]) == comparer {
			filteredList = append(filteredList, line)
		}
	}
	return filteredLines(filteredList, i+1, isCommon)
}
func part2(lines []string){
	oxygen := filteredLines(lines, 0, true)
	scrubber := filteredLines(lines, 0, false)
	oxygenInt, _ := strconv.ParseInt(oxygen,2,64)
	scrubberInt, _ := strconv.ParseInt(scrubber,2,64)
	fmt.Println(int(oxygenInt) * int(scrubberInt))
}
func main() {
	lines, _:= utils.ReadStringArrayFromFile("input.txt")
	part1(lines)
	part2(lines)
}
