package main

type Product struct {
	Name  string
	Price float64
}

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
}

func (f *Filter) FilterProducts(products []Product, filter func(Product) bool) []Product {
	var result []Product
	for _, p := range products {
		if filter(p) {
			result = append(result, p)
		}
	}
	return result
}

func main() {
	var p Product = Product{Name: "Apple", Price: 1.0}
	var f Filter = Filter{}
	f.FilterProducts([]Product{p}, func(p Product) bool {
		return p.Price > 1.0
	}).Print(" - %s greater than 1.0", "Price")
}
