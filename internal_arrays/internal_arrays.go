package main

import "fmt"

func main() {

  slice := make([]float32, 10, 11)
  fmt.Println(slice)
  fmt.Println(len(slice))
  fmt.Println(cap(slice))

  slice = append(slice, 5)
  fmt.Println(slice)
  fmt.Println(len(slice))
  fmt.Println(cap(slice))

  slice = append(slice, 5)
  fmt.Println(slice)
  fmt.Println(len(slice))
  fmt.Println(cap(slice))

  slice1 := make([]float32, 10)
  fmt.Println(slice1)
  fmt.Println(len(slice1))
  fmt.Println(cap(slice1))
  
}