package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := float64(1)
  for i, eps := 0, 1.0; math.Abs(eps) > 1e-2; i++ {
    eps = (z*z - x) / 2*z
    z = z - eps
    fmt.Printf("iteration %d, error %f\n", i, math.Abs(eps))
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(math.Sqrt(2))
}

