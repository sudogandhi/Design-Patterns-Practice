package main

// Composite design pattern is used when we can arrage the elements in the form of an hierarchy.
// Each of the class implements the same interface
// There can be multiple examples of such kind of hierarchies :-
// 1. Box and Product - A Box can contain multiple boxes and multipe packages and each package will represent the leaf node.
// 2. Directory and files
// 3. Branches and Leaves
// 4. Teacher and Student

// Let's implement a scenario of Box and Package

import "fmt"

type Component interface {
	computePrice() float64
}

type Box struct {
	components []Component
}

func (box *Box) computePrice() float64 {
	var price float64
	// each box will add up a $2 price as packaging charges
	price = 2

	for each := range box.components {
		price += box.components[each].computePrice()
	}
	return price
}

func (box *Box) add(component Component) {
	box.components = append(box.components, component)
}

type Product struct {
	price float64
}

func (product *Product) computePrice() float64 {
	return product.price
}

func (product *Product) setPrice(price float64) {
	product.price = price
}

func main() {
	fmt.Println("Implementing composite pattern")

	box1 := Box{}
	box2 := Box{}
	box1.add(&box2)
	box3 := Box{}
	box1.add(&box3)

	p1 := Product{}
	p1.setPrice(10)
	box1.add(&p1)

	p2 := Product{}
	p2.setPrice(20)
	box2.add(&p2)

	p3 := Product{}
	p3.setPrice(30)
	box2.add(&p3)

	p4 := Product{}
	p4.setPrice(40)
	box3.add(&p4)

	fmt.Printf("Price of Box 1 : %0.2f", box1.computePrice())
}
