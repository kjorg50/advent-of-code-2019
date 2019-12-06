package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  if len(os.Args) != 4 {
    fmt.Println("Usage: go run dayXX.go -- part[1|2] <filename>")
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
  for scanner.Scan() {
    // get line of text with scanner.Text()

  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

}

func part2(scanner *bufio.Scanner) {
  for scanner.Scan() {
    // get line of text with scanner.Text()

  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  
}
