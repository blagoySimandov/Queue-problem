package main

import "fmt"

type (
	t int

	queueInterface interface {
		push(v t)
		pop() t
		getMax() t
		printMe()
	}

	queue struct {
		nodes *node
	}

	node struct {
		value t
		prev  *node
		max   t
	}
)

func (q *queue) push(v t) {
	n := node{
		value: v,
		prev:  q.nodes, // nil
	}

	q.nodes = &n

	if v > n.max {
		n.max = v
	}
}

func (q *queue) pop() t {
	n := q.nodes.prev
	oldValue := q.nodes.value
	q.nodes = n
	return oldValue
}

func (q *queue) getMax() t {
	if q.nodes != nil {
		return q.nodes.max
	}
	return -1
}

func (q queue) printMe() {
	fmt.Print("[")
	for n := q.nodes; n != nil; n = n.prev {
		fmt.Printf("%v, ", n.value)
	}
	fmt.Println("]")
}

func main() {
	var q queue
	q.push(10)
	q.push(20)
	q.push(30)
	q.printMe()
}
