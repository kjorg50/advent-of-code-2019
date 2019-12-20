package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func check(e error) {
  if e != nil {
    log.Fatal(e)
    panic(e)
  }
}

// Abs returns the absolute value of x.
func Abs(x int64) int64 {
  if x < 0 {
    return -x
  }
  return x
}

func main() {
  if len(os.Args) != 4 {
    fmt.Println("Usage: go run day03.go -- part[1|2] <filename>")
    os.Exit(1)
  }
  part := os.Args[2]
  filename := os.Args[3]

  file, err := os.Open(filename)
  check(err)

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

type Point struct{
  x, y int64
}


var (
  DX = map[string]int {"U": 0, "D": 0, "L": -1, "R": 1}
  DY = map[string]int {"U": 1, "D": -1, "L": 0, "R": 0}
)

func insertPoints(path string) map[Point]bool {
  result := map[Point]bool {}
  s := strings.Split(path, ",")
  x, y := 0, 0
  for _, cmd := range s {
    // get the direction and the value of the command
    dir := string(cmd[0])
    rest := cmd[1:len(cmd)]
    num, err := strconv.Atoi(rest)
    check(err)

    // add all the points on this line to the path
    for i := 0; i < num; i++ {
      x += DX[dir]
      y += DY[dir]
      result[Point{int64(x),int64(y)}] = true
    }
  }

  return result
}

func part1(scanner *bufio.Scanner) {
  var pathA = map[Point]bool {}
  var pathB = map[Point]bool {}
  i := 0

  // read the data into the paths (hack the for loop for two lines)
  for scanner.Scan() {
    switch i {
    case 0:
      pathA = insertPoints(scanner.Text())
    case 1:
      pathB = insertPoints(scanner.Text())
    }
    i++
  }

  // find the Manhattan distances of the points
  var distances []int64
  for p := range pathA {
    _, ok := pathB[p]
    if ok {
      distances = append(distances, Abs(p.x) + Abs(p.y))
    }
  }

  // get the min distance
  min := distances[0]
  for _, v := range distances {
    if v < min {
      min = v
    }
  }

  err := scanner.Err()
  check(err)

  fmt.Println("Min is", min)
}

func part2(scanner *bufio.Scanner) {
  for scanner.Scan() {
    // get line of text with scanner.Text()

  }

  err := scanner.Err()
  check(err)
}
