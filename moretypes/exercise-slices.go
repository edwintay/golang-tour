package main

import (
  "golang.org/x/tour/pic"
  "math"
)

func createPixel(row, col int) uint8 {
  // return uint8((row+col)/2)
  // return uint8(row*col)
  return uint8(math.Pow(float64(col), float64(row)))
}

func Pic(dx, dy int) [][]uint8 {
  output := make([][]uint8, dy)
  for y := 0; y < dy; y++ {
    row := make([]uint8, dx)
    for x := 0; x < dx; x++ {
      row[x] = createPixel(y, x)
    }
    output[y] = row
  }
  return output
}

func main() {
  pic.Show(Pic)
}

