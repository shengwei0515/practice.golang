package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkRecursive(t.Left, ch)
	ch <- t.Value
	// fmt.Println(t.Value)
	WalkRecursive(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chanT1 := make(chan int)
	chanT2 := make(chan int)

	go Walk(t1, chanT1)
	go Walk(t2, chanT2)

	for {
		v1, ok1 := <-chanT1
		v2, ok2 := <-chanT2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 || !ok2 {
			break
		}
	}

	return true
}

func main() {
	channel := make(chan int)
	go Walk(tree.New(10), channel)
	func() {
		for v := range channel {
			fmt.Println(v)
		}
	}()
	fmt.Println(Same(tree.New(10), tree.New(20)))
	fmt.Println(Same(tree.New(10), tree.New(10)))
	fmt.Println(Same(tree.New(20), tree.New(10)))
}
