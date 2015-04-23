package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  wordMap := make(map[string]int)
  for _, word := range strings.Fields(s) {
    wordMap[word] = wordMap[word] + 1
  }
  return wordMap
}

func main() {
  wc.Test(WordCount)
}

