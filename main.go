package main

import (
	"fmt"
	"math"
	"time"
)

func linearRNG(x0, a, b, m, n int) []int {
  /*
  x0: Seed
  a: Multiplier
  b: Increment
  m: Modulus
  n: Amount of numbers
  */
  results := []int{}
  // Checking necessary conditions
  if !(m > 0) {
    panic("The modulus must be positive")
  }
  if !(0 < a && a < m) {
    panic("The multiplier must be positive and less than modulus")
  }
  if !(0 <= b && b < m) {
    panic("The multiplier must be non-negative and less than modulus")
  }
  for i := 1; i <= n; i++ {
    pos := false
    for !pos {
      x0 = (a * x0 + b) % m
      if x0 >= 0 {
        pos = true
        results = append(results, x0)
      }
    }
  }
  return results
}

func main() {
  fmt.Println("====================")
  x0 := int(math.Round(float64(time.Now().UnixNano())))
  var a, b, m, n int
  fmt.Println("Enter multiplier, increment, modulus and amount of numbers to generate (in that order):")
  fmt.Scan(&a, &b, &m, &n)
  fmt.Println("Here is your list of numbers:")
  fmt.Println(linearRNG(x0, a, b, m, n))
}
