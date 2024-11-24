package main

import "fmt"

type Queue struct {
	items []int
	size  int
	front int
	rear  int
}

func InitialiseQueue(size int) *Queue {
	q := &Queue{}
	q.items = make([]int, size)
	q.size = size
	q.front = -1
	q.rear = -1
	return q
}

func (q *Queue) Enqueue(n int) {
	if q.rear == q.size-1 {
		fmt.Println("Queue is full")
	} else {
		if q.front == -1 {
			q.front = 0
			q.rear = 0
		} else {
			q.rear = q.rear + 1
		}
		q.items[q.rear] = n
	}
}

func (q *Queue) Dequeue() {
	if q.front == -1 || q.front > q.rear {
		fmt.Println("Queue is empty")
	} else {
		temp := q.items[q.front]
		if q.front == q.rear {
			q.front = -1
		} else {
			q.front = q.front + 1
		}
		fmt.Printf("%d is DeQueued\n", temp)
	}
}

func (q *Queue) ListQueue() {
	if q.front == -1 {
		fmt.Println("Queue is empty")
	} else {
		fmt.Printf("Queue items: ")
		for i := q.front; i <= q.rear; i++ {
			fmt.Printf("%d \n", q.items[i])
		}
	}
}
func main() {
	fmt.Println("Implementing Queue Data Structure")

	q := InitialiseQueue(5)
	q.Enqueue(10)
	q.Enqueue(16)
	q.Enqueue(11)
	q.ListQueue()
	q.Dequeue()
	q.ListQueue()
}
