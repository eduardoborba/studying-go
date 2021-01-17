package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walker(t, ch)

	close(ch)
}

func walker(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walker(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walker(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 {
			return false
		}

		if !ok1 || !ok2 {
			break
		}
	}

	return true
}

func main() {
	fmt.Println("1 and 1 same: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("1 and 2 same: ", Same(tree.New(1), tree.New(2)))
}
