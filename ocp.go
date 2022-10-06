package main

// open closed principle
type Product struct {
	productName string
	color       Color
	size        Size
}

// a type oc Color
type Color int

const (
	Red Color = iota
	Blue
	Green
)

// a type of size
type Size int

const (
	Large Size = iota
	Medium
	Small
)

// declare a filter type
type Filter struct{} // this will be the reciever

/* implementing the open closed solid principle */
func (f *Filter) FilterByColor(p []Product, color Color) []*Product {

	reqProduct := make([]*Product, 0)

	//iterate over the color and return any product that matches the specification
	for i, val := range p {

		if color == val.color {
			reqProduct = append(reqProduct, &p[i])
		}
	}

	return reqProduct
}

// create a specification
type Specification interface {
	isSatisfied(p *Product) bool
}
type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

type AndSpecification struct {
	first, second Specification
}

func (a *AndSpecification) isSatisfied(p *Product) bool {
	return a.first.isSatisfied(p) && a.second.isSatisfied(p)
}

func (c *ColorSpecification) isSatisfied(p *Product) bool {
	return c.color == p.color
}

func (s *SizeSpecification) isSatisfied(p *Product) bool {
	return s.size == p.size
}

type BetterFilter struct {
}

func (b *BetterFilter) AdvancedFilter(p []Product, s Specification) []*Product {
	value := make([]*Product, 0)

	for i, val := range p {
		if s.isSatisfied(&val) {
			value = append(value, &p[i])
		}
	}

	return value
}


// f := Filter{} // this way we can have access to the methods associated with the struct type

// 	//define a products
// 	greenProduct := Product{"Green Ball", Green, Small}
// 	blueProduct := Product{"Blue Ball", Blue, Large}
// 	redProduct := Product{"Red Ball", Red, Medium}

// 	vProduct := []Product{greenProduct, blueProduct, redProduct}

// 	fmt.Println("This is the old method:\n")
// 	for _, v := range f.FilterByColor(vProduct, Red) {
// 		fmt.Printf("- The name of the Product is %s and the color is %+v", v.productName, v.color)
// 	}

// 	bf := BetterFilter{}
// 	colorSpec := &ColorSpecification{Blue}
// 	fmt.Println("\nThis is the new method")
// 	for _, v := range bf.AdvancedFilter(vProduct, colorSpec) {
// 		fmt.Printf("- The name of the Product is %s and the color is %+v", v.productName, v.color)
// 	}

// 	//check for two criterias
// 	fmt.Println("\n The aggregated specification")
// 	largeProduct := &SizeSpecification{Large}
// 	andSpec := &AndSpecification{colorSpec, largeProduct}
// 	for _, v := range bf.AdvancedFilter(vProduct, andSpec) {
// 		fmt.Printf("- The name of the Product is %s and the color is %+v", v.productName, v.color)
// 	}