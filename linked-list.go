package main

import "fmt"


type node struct {
	val int
	next *node
}



// walk to end of the linked list and add a new node with x.
func appnd(curr *node, x int) {
	// can't append to a nil node
	if curr == nil {
		return
	} else {
		for (*curr).next != nil {
			curr = (*curr).next
		}
		tmp := node{val: x}
		(*curr).next = &tmp
	}
}


// print the linked list!
func prnt(head *node) {
	tmp := head
	for tmp != nil {
		fmt.Println((*tmp).val)
		tmp = (*tmp).next
	}
}

func main() {
	tmp := node{val: 5}
	appnd(&tmp, 6)
	appnd(&tmp, 7)
	prnt(&tmp)
	
}
