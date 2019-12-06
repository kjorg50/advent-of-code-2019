package main

import (
	"fmt"
	"strconv"

	//"go/scanner"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run dayXX.go -- part[1|2] <filename>")
		os.Exit(1)
	}
	part := os.Args[2]
	filename := os.Args[3]

	b, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(b),",")
	nums := make([]int, 0, len(lines))

	for _, v := range lines {
		if len(v) == 0 { continue }

		n, err := strconv.Atoi(v)
		check(err)
		nums = append(nums, n)
	}

	switch part {
	case "part1":
		fmt.Println(part1(nums))
	case "part2":
		noun, verb := part2(nums)
		fmt.Println(100 * noun + verb)
	default:
		fmt.Println("You must select 'part1' or 'part2'")
		os.Exit(1)
	}
}

func part1(numbers []int) int {
	return runLoop(12, 2, numbers)
}

func part2(numbers []int) (int, int) {
	goal := 19690720

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if runLoop(i, j, numbers) == goal {
				fmt.Printf("noun is %v and verb is %v\n", i, j)
				return i, j
			}
		}
	}
	// bogus error return val
	return -1, -1
}

func runLoop(noun int, verb int, numbers []int) int {
	var result = make([]int, len(numbers))
	copy(result, numbers)

	result[1] = noun
	result[2] = verb

	for i := 0; i < len(result); i += 4 {
		switch result[i] {
		case 1:
			result[ result[i+3] ] = result[ result[i+1] ] + result[ result[i+2] ]
		case 2:
			result[ result[i+3] ] = result[ result[i+1] ] * result[ result[i+2] ]
		case 99:
			return result[0]
		}
	}
	return result[0]
}
