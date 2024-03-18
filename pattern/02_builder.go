package main

import "fmt"

type Product struct{
	part1 string
	part2 string
	part3 string
}

type Builder interface{
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetProduct() *Product
}

func NewConcreteBuilder() *ConcreteBuilder{
	return &ConcreteBuilder{
		product: &Product{},
	}
}
type ConcreteBuilder struct{
	product *Product
}

func (c *ConcreteBuilder) BuildPart1(){
	c.product.part1 = "part1"
}

func (c *ConcreteBuilder) BuildPart2(){
	c.product.part2 = "part2"
}

func (c *ConcreteBuilder) BuildPart3(){
	c.product.part3 = "part3"
}

func (c *ConcreteBuilder) GetProduct() *Product{
	return c.product
}

type Director struct{
	builder Builder
}

func NewDirector(builder Builder) *Director{
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct(){
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
}

func main(){
	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetProduct()
	fmt.Println(product.part1, product.part2, product.part3)
}

