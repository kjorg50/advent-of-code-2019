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
		//part2(nums)
	default:
		fmt.Println("You must select 'part1' or 'part2'")
		os.Exit(1)
	}
}

func part1(numbers []int) int {
	var result = make([]int, len(numbers))
	copy(result, numbers)

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

//func part2(numbers []int) {
//	for scanner.Scan() {
//		// get line of text with scanner.Text()
//
//	}
//
//	err := scanner.Err()
//	check(err)
//}
