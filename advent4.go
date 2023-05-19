package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const file = "inputs.txt"

func main() {
	// Open the file
	file, _ := os.Open(file)
	scanner := bufio.NewScanner(file)
	sumContains := 0
	sumOverlaps := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitComma := strings.Split(line, ",")
		splitDash1 := strings.Split(splitComma[0], "-")
		splitDash2 := strings.Split(splitComma[1], "-")

		assignment1Start, _ := strconv.Atoi(splitDash1[0])
		assignment1End, _ := strconv.Atoi(splitDash1[1])
		var assignment1 []int

		for i := assignment1Start; i < assignment1End+1; i++ {
			assignment1 = append(assignment1, i)
		}

		assignment2Start, _ := strconv.Atoi(splitDash2[0])
		assignment2End, _ := strconv.Atoi(splitDash2[1])
		var assignment2 []int

		for i := assignment2Start; i < assignment2End+1; i++ {
			assignment2 = append(assignment2, i)
		}

		if checkAssignmentContains(assignment1, assignment2) || checkAssignmentContains(assignment2, assignment1) {
			sumContains++
		}

		if checkOverlaps(assignment1, assignment2) {
			sumOverlaps++
		}

	}
	log.Println(sumContains)
	log.Println(sumOverlaps)
}

func checkAssignmentContains(assignment1 []int, assignment2 []int) bool {
	for _, val1 := range assignment1 {
		exists := false
		for _, val2 := range assignment2 {
			if val1 == val2 {
				exists = true
			}
		}
		if !exists {
			return false
		}
	}
	return true
}

func checkOverlaps(assignment1 []int, assignment2 []int) bool {
	for _, val1 := range assignment1 {
		for _, val2 := range assignment2 {
			if val1 == val2 {
				return true
			}
		}
	}
	return false
}
