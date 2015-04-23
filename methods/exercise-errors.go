package main

import (
  "fmt"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %f", float64(err))
}

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }
  return 0, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}

