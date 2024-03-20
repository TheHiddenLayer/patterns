package main

import "fmt"

type IProduct interface {
	Operation()
}

type Product struct {
	Name string
}

func (p *Product) Operation() {
	fmt.Println("Product Name:", p.Name)
}

type ProductA struct {
	Name string
	Product
}

type ProductB struct {
	Name string
	Product
}

func NewProductA() IProduct {
	return &ProductA{Name: "A"}
}

func NewProductB() IProduct {
	return &ProductB{Name: "B"}
}

func GetProduct(productType string) (IProduct, error) {
	switch productType {
	case "A":
		return NewProductA(), nil
	case "B":
		return NewProductB(), nil
	default:
		return nil, fmt.Errorf("invalid product type")
	}
}
