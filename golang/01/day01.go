package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) != 4 {
    fmt.Println("Usage: go run day01.go -- part[1|2] <filename>")
    os.Exit(1)
  }
  part := os.Args[2]
  filename := os.Args[3]

  file, err := os.Open(filename)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)

  switch part {
  case "part1":
    part1(scanner)
  case "part2":
    part2(scanner)
  default:
    fmt.Println("You must select 'part1' or 'part2'")
    os.Exit(1)
  }
}

func part1(scanner *bufio.Scanner) {
  result := 0
  for scanner.Scan() {
    if i, err := strconv.Atoi(scanner.Text()); err == nil {
      mass := i / 3 - 2
      result += mass
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(result)
}

func part2(scanner *bufio.Scanner) {
  result := 0
  for scanner.Scan() {
    if i, err := strconv.Atoi(scanner.Text()); err == nil {
      result += calcRemainingFuel(i)
    }
  }
  fmt.Println(result)
}

func calcRemainingFuel(x int) int {
  neededFuel := x / 3 - 2
  if (neededFuel) < 1 {
    return 0
  } else {
    return neededFuel + calcRemainingFuel(neededFuel)
  }
}
