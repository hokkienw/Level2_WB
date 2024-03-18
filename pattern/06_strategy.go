package main

import "fmt"

type strategy interface{
	Execute(int, int) int
}
type add struct{}

type sub struct{}

func (a add) Execute(a1, a2 int) int {
	return a1 + a2
}

func (s sub) Execute(s1, s2 int) int {
	return s1 - s2
}
type Context struct{
	strategy
}

func (c *Context) SetStrategy(s strategy){
	c.strategy = s
}

func (c *Context) Execute(a1, a2 int) int {
	return c.strategy.Execute(a1, a2)
}

func NewContext(strategy strategy) *Context{
	return &Context{
		strategy: strategy,
	}
}

func main(){
	add := NewContext(add{})
	fmt.Println(add.Execute(1, 2))

	sub := NewContext(sub{})
	fmt.Println(sub.Execute(1, 2))
}


