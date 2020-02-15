package structure

type CirCleQueue struct {
	data []int
	head int
	tail int
	size int
}

func NewCirCleQueue(size int) *CirCleQueue {
	return &CirCleQueue{
		data: make([]int, size+1),
		head: -1,
		tail: -1,
		size: size,
	}
}

func (queue *CirCleQueue) Enqueue(e int) bool {
	if queue.IsFull() {
		return false
	}
	if queue.IsEmpty() {
		queue.head = 0
	}

	queue.tail = (queue.tail + 1) % queue.size
	queue.data[queue.tail] = e
	return true
}

func (queue *CirCleQueue) Dequeue() (int, bool) {
	if queue.IsEmpty() {
		return -1, false
	}
	e := queue.data[queue.tail]
	if queue.head == queue.tail {
		queue.head = -1
		queue.tail = -1
	} else {
		queue.head = (queue.head + 1) % queue.size
	}

	return e, true
}

func (queue *CirCleQueue) Front() int {
	if queue.IsEmpty() {
		return -1
	}

	return queue.data[queue.head]
}

func (queue *CirCleQueue) Rear() int {
	if queue.IsEmpty() {
		return -1
	}

	return queue.data[queue.tail]
}

func (queue *CirCleQueue) IsFull() bool {
	return (queue.tail+1)%queue.size == queue.head
}

func (queue *CirCleQueue) IsEmpty() bool {
	return queue.head == -1
}
