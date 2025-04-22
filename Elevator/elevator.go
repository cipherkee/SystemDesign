package main

import (
	"fmt"
)

type Direction int

var (
	UpDirection   Direction = 1
	DownDirection Direction = -1
)

type InRequestNode struct {
	floor int
	right *InRequestNode
	left  *InRequestNode
}

func NewInRequestNode(floor int) *InRequestNode {
	return &InRequestNode{
		floor: floor,
	}
}

type InRequests struct {
	head   *InRequestNode
	tail   *InRequestNode
	reqMap map[int]*InRequestNode
}

func (r *InRequests) AddRequest(floor int) {
	_, ok := r.reqMap[floor]
	if ok {
		return
	}
	node := NewInRequestNode(floor)
	r.reqMap[floor] = node

	// adding to ll
	if r.tail == nil {
		r.tail = node
		r.head = node
		return
	}

	r.tail.right = node
	node.left = r.tail
	r.tail = node
}

type Elevator struct {
	dispatcher *Dispatcher
	currFloor  int
	requests   *InRequests
	direction  Direction
}

func (e *Elevator) getFromInnerRequest() int {
	if e.requests.head == nil {
		return -1
	}
	head := e.requests.head
	e.requests.head = head.right
	head.right.left = nil
	return head.floor
}

func (e *Elevator) Move() {
	targetFloor := e.getFromInnerRequest()

	if targetFloor == -1 {
		targetFloor = e.dispatcher.dispatchNextRequest()
	}

	// change the direction of elevator if target is not in the direction
	if (e.currFloor < targetFloor && e.direction == DownDirection) || (e.currFloor > targetFloor && e.direction == UpDirection) {
		e.direction = Direction(e.direction * -1)
	}

	for e.currFloor != targetFloor {
		if e.dispatcher.hasRequest(e.currFloor, e.direction) {
			fmt.Printf("Landing in floor %v\n", e.currFloor)
		}
		e.currFloor += int(e.direction)
	}
}
