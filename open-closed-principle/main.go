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

type Specification interface {
	IsSatisfied(p *Product) bool
}

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

type BetterFilter struct {

}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
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



	/// New filter
	fmt.Println("\n\nNew implementation \n")
	bf := BetterFilter{}
	greenSpecification := ColorSpecification{green}
	for _, v := range bf.Filter(products, greenSpecification) {
		fmt.Printf("    %s is green ", v.name)
	}

	sizeSpecification := SizeSpecification{large}
	for _, v := range bf.Filter(products, sizeSpecification) {
		fmt.Printf("    %s is large ", v.name)
	}
}