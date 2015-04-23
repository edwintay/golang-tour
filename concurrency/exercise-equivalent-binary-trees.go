package main

import (
  "fmt"
  "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  walk(t, ch)
  close(ch)
}

func walk(t *tree.Tree, ch chan int) {
  if t == nil {
    return
  }
  walk(t.Left, ch)
  ch <- t.Value
  walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  s1, ch1 := make([]int, 10), make(chan int)
  s2, ch2 := make([]int, 10), make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
 
  for {
    select {
      case v, ok := <-ch1:
        if ok {
          s1 = append(s1, v)
        } else {
          ch1 = nil
        }
      case v, ok := <-ch2:
        if ok {
          s2 = append(s2, v)
        } else {
          ch2 = nil
        }
    }

    if ch1 == nil && ch2 == nil {
      break
    }
  }

  fmt.Println(s1)
  fmt.Println(s2)
  return same(s1, s2)
}

func same(s1, s2 []int) bool {
  var i int
  for i = 0; i<len(s1) && i<len(s2); i++ {
    if s1[i] != s2[i] {
      return false
    }
  }
  if i < len(s1) || i < len(s2) {
    return false // unequal lengths
  }
  return true
}

func main() {
  fmt.Println(Same(tree.New(1), tree.New(1)))
  fmt.Println(Same(tree.New(1), tree.New(2)))
}

