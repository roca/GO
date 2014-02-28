// tour71
package main

import "code.google.com/p/go-tour/tree"
import "fmt"

func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t.Left != nil {
			walker(t.Left)
		}
		ch <- t.Value
		if t.Right != nil {
			walker(t.Right)
		}
	}
	walker(t)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	chana := make(chan int)
	chanb := make(chan int)

	go Walk(t1, chana)
	go Walk(t2, chanb)

	for {
		n1, ok1 := <-chana
		n2, ok2 := <-chanb
		fmt.Println("values: ", n1, n2)
		if n1 != n2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {

	t1 := tree.Tree{
		&tree.Tree{
			&tree.Tree{
				nil,
				1,
				nil,
			},
			1,
			&tree.Tree{
				nil,
				2,
				nil,
			},
		},
		3,
		&tree.Tree{
			&tree.Tree{
				nil,
				5,
				nil,
			},
			8,
			&tree.Tree{
				nil,
				13,
				nil,
			},
		},
	}
	t2 := tree.Tree{
		&tree.Tree{
			&tree.Tree{
				&tree.Tree{nil, 1, nil},
				1,
				&tree.Tree{nil, 2, nil},
			},
			3,
			&tree.Tree{nil, 5, nil},
		},
		8,
		&tree.Tree{nil, 13, nil},
	}

	fmt.Println("Are they the same ?", Same(&t1, &t2))
}
