package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(b []byte) (int, error) {
  copy(b, []byte{'A'})
  return 1, nil
}

func main() {
  reader.Validate(MyReader{})
}

