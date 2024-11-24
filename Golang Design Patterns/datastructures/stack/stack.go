package main

import (
	"errors"
	"fmt"
)

const (
	INITIAL_STACK_SIZE    = 10
	STACK_FILL_PERCENTAGE = .7
)

var ErrEmptyStack = errors.New("stack is empty")
var ErrInvalidChoice = errors.New("invalid choice")

type Stack struct {
	elements []int
	top      int
	size     int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]int, INITIAL_STACK_SIZE),
		top:      -1,
		size:     INITIAL_STACK_SIZE,
	}
}

func (s *Stack) Push(n int) {
	s.top++
	s.elements[s.top] = n
	s.Resize()
}

func (s *Stack) Pop() (int, error) {
	if s.top < 0 {
		return 0, ErrEmptyStack
	} else {
		n := s.elements[s.top]
		s.top--
		return n, nil
	}
}

func (s *Stack) PrintStack() {
	fmt.Println()
	for i := 0; i <= s.top; i++ {
		fmt.Printf("%d ", s.elements[i])
	}
	fmt.Println()
}

func (s *Stack) Resize() {
	if float64(s.top) >= float64(s.size)*STACK_FILL_PERCENTAGE {
		newSize := s.size * 2
		s.elements = append(s.elements, make([]int, newSize-s.size)...)
		s.size = newSize
	}
}

func (s *Stack) PrintStackSize() {
	fmt.Printf("Stack size: %d\n", s.size)
}

func (s *Stack) PrintTopValue() {
	fmt.Printf("Top value: %d\n", s.top)
}

func (s *Stack) Peep() error {
	if s.top == -1 {
		return ErrEmptyStack
	} else {
		fmt.Printf("Top Element :: %d\n", s.elements[s.top])
		return nil
	}
}

func PrintOptions() {
	fmt.Println("1. Push")
	fmt.Println("2. Pop")
	fmt.Println("3. Peep")
	fmt.Println("4. Print Stack")
	fmt.Println("5. Stack Size")
	fmt.Println("6. Top Value")
	fmt.Println("7. Exit")
	fmt.Print("\nEnter your choice :: ")
}

func main() {
	fmt.Println("Implementing stacks")
	var stack *Stack = NewStack()
	var ch int
	for {
		PrintOptions()
		fmt.Scan(&ch)
		if ch == 1 {
			fmt.Print("Enter the element to push :: ")
			var n int
			fmt.Scan(&n)
			stack.Push(n)
		} else if ch == 2 {
			n, err := stack.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%d element is popped", n)
			}
		} else if ch == 3 {
			stack.Peep()
		} else if ch == 4 {
			stack.PrintStack()
		} else if ch == 5 {
			stack.PrintStackSize()
		} else if ch == 6 {
			stack.PrintTopValue()
		} else if ch == 7 {
			break
		} else {
			fmt.Println(ErrInvalidChoice)
		}
	}
}
