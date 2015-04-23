package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
  n, err:= r.r.Read(b)
  if err != nil {
    return 0, err
  }

  result := b[:n]
  for i, c := range result {
    if 'A' <= c && c <= 'Z' {
      result[i] = rot13Upper(c)
    } else if 'a' <= c && c <= 'z' {
      result[i] = rot13Lower(c)
    }
  }
  return n, nil
}

func rot13Upper(b byte) (byte) {
  return rot13(b, 'A', 'Z')
}

func rot13Lower(b byte) (byte) {
  return rot13(b, 'a', 'z')
}

func rot13(b, min, max byte) (byte) {
  diff := int(b) - int(min)
  maxdiff := int(max) - int(min)
  roted := int(min) + (diff+13) % (maxdiff+1)
  return byte(roted)
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}

