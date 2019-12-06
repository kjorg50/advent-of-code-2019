package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
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

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      // do stuff
      fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
}
