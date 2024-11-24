package main

import "fmt"

type Set struct {
	data map[int]bool
}

func (set *Set) New() {
	set.data = make(map[int]bool)
}

func (set *Set) AddElement(val int) {
	if !set.ContainsElement(val) {
		set.data[val] = true
	}
}

func (set *Set) ContainsElement(val int) bool {
	_, ok := set.data[val]
	return ok
}

func (set *Set) DeleteElement(val int) {
	delete(set.data, val)
}

func main() {
	fmt.Println("Implementing Set Data Structure")
}
