package main

import "fmt"

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

type Queue []*Order

func (order *Order) New(p int, q int, prod string, cust string) {
	order.priority = p
	order.quantity = q
	order.product = prod
	order.customerName = cust
}

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false
		// performing insertion sort
		for i, o := range *queue {
			if o.priority > order.priority {
				// *queue[:i] will give us all the elements from 0 till i-1
				// Queue{order} will create an array of orders and then we will append the rest of the elements from i till last
				// Then append the newly formed array with the starting i elements
				// [0,1,2,3......i-1] order [ i, i+1, i+2 ....]
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}

		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func (queue *Queue) PrintElements() {
	fmt.Println("Customer\t\tProduct\t\tQuantity\t\tPriority")
	for _, order := range *queue {
		fmt.Printf("%s\t\t%s\t\t%d\t\t%d\n", order.customerName, order.product, order.quantity, order.priority)
	}
}
func main() {
	fmt.Println("Implementing Priority Queue")
	q := Queue{}
	q.Add(&Order{1, 10, "ABC", "QWERTY"})
	q.Add(&Order{15, 10, "ABC", "QWERTY"})
	q.Add(&Order{3, 10, "ABC", "QWERTY"})
	q.Add(&Order{4, 10, "ABC", "QWERTY"})
	q.Add(&Order{5, 10, "ABC", "QWERTY"})
	q.Add(&Order{2, 10, "ABC", "QWERTY"})
	q.Add(&Order{12, 10, "ABC", "QWERTY"})

	q.PrintElements()
}
