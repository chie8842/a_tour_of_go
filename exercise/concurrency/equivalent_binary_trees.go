package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	walk := func(t *tree.Tree){}
	walk = func(t *tree.Tree) {
		if t.Left != nil {
			//fmt.Println(t.Left)
			walk(t.Left)
		}
		//fmt.Println(t.Value,)
		ch <- t.Value
		if t.Right != nil {
			walk(t.Right)
		}
	}
	walk(t)
	close(ch)
}


// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	fmt.Println(t1, t2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		j, ok := <- ch2
		if ! ok || i != j {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
