package main

import "fmt"

type Color int
const (
	red Color = iota
	green
	blue
)

type Size int
const (
	small Size = iota
	medium
	large
)

type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

type Product struct {
	name string
	color Color
	size Size
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Println("Green products (old):\n")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf("    %s is green ", v.name)
	}

	for _, v := range f.FilterBySize(products, large) {
		fmt.Printf("    %s is large ", v.name)
	}
}