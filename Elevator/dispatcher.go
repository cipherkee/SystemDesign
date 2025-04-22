package main

/*
	Dispatcher needs to maintain a queue of external requests.
	It should also be able to delete requests, if the elevator is able to address it on the way.
	Its a combination of FIFO, while also serving requests that are on the way.
*/

type RequestNode struct {
	Right *RequestNode
	Left  *RequestNode
	floor int
}

type Requests struct {
	head  *RequestNode
	tail  *RequestNode
	index map[int]*RequestNode // map of floor no to the request
}

func (r *Requests) dispatchNextRequest() int {
	if r.head == nil {
		return -1
	}
	req := r.head
	r.head = r.head.Right
	r.head.Left = nil
	delete(r.index, req.floor)

	return req.floor
}

type Dispatcher struct {
	upRequest   *Requests
	downRequest *Requests
}

func (d *Dispatcher) hasRequest(floor int, direction Direction) bool {
	// TODO: If has request, then delete request and return

	return false
}

// First return if up request present, else down request
func (d *Dispatcher) dispatchNextRequest() int {
	floor := d.upRequest.dispatchNextRequest()
	if floor != -1 {
		return floor
	}
	floor = d.downRequest.dispatchNextRequest()
	return floor
}
