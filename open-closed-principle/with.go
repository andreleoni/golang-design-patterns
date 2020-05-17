package main

import "fmt"

// Product attributes to be filtered

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

type Product struct {
	name string
	color Color
	size Size
}

// Filter structure to filter based on implementation specifications

type Specification interface {
	IsSatisfied(p *Product) bool
}

type Filter struct {}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func (f *Filter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

// Implementation of filters to color and size through the interfaces

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	f := Filter{}
	greenSpecification := ColorSpecification{green}
	for _, v := range f.Filter(products, greenSpecification) {
		fmt.Printf("    %s is green ", v.name)
	}

	sizeSpecification := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpecification, sizeSpecification}
	for _, v := range f.Filter(products, lgSpec) {
		fmt.Printf("    %s is large and green ", v.name)
	}
}