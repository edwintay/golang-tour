package main

import (
  "golang.org/x/tour/pic"
  "image"
  "image/color"
)

type Image struct{
  x, y, w, h int
}

func (img Image) ColorModel() (color.Model) {
  return color.RGBAModel
}

func (img Image) Bounds() (image.Rectangle) {
  return image.Rect(img.x, img.y, img.w, img.h)
}

func (img Image) At(x, y int) (color.Color) {
  mx, my := x-img.x, y-img.y
  v := uint8((mx+my)/4)
  return color.RGBA{v, v, 255, 255}
}

func main() {
  m := Image{x: 128, y: 128, w: 512, h: 512}
  pic.ShowImage(m)
}

