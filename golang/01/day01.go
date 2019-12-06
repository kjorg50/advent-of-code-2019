package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) != 3 {
    fmt.Println("Usage: go run day01.go -- <filename>")
  }
  filename := os.Args[2]

  file, err := os.Open(filename)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  result := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if i, err := strconv.Atoi(scanner.Text()); err == nil {
      mass := (i / 3) - 2
      result += mass
    }
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  fmt.Println(result)
}
