package main

import (
    "fmt"
	"os"
 	"strings"
	"strconv"
	"bufio"
)


func Part1(sc *bufio.Scanner) int{
	var depth, horizontal int
	for sc.Scan(){
		movement := strings.Fields(sc.Text())
		increment, _ := strconv.Atoi(movement[1])
		switch(movement[0]){
		case "forward":
			horizontal += increment
		case "down":
			depth += increment
		case "up":
			depth -= increment
		}
	}
	return depth * horizontal
}
func Part2(sc *bufio.Scanner) int{
	var depth, horizontal, aim int
	for sc.Scan(){
		movement := strings.Fields(sc.Text())
		increment, _ := strconv.Atoi(movement[1])
		switch(movement[0]){
		case "forward":
			horizontal += increment
			depth += (aim * increment)
		case "down":
			aim += increment
		case "up":
			aim -= increment
		}
	}
	return depth * horizontal
}

func main() {
	input1, _ := os.Open("input.txt")
	input2, _ := os.Open("input.txt")
	sc1 := bufio.NewScanner(input1)
	sc2 := bufio.NewScanner(input2)
	fmt.Println("Part 1: " , Part1(sc1))
	fmt.Println("Part 2: " , Part2(sc2))
	
}
