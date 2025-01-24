package main

import (
	"fmt"
)

func main() {
	i := NewItem(10, 20)
	p1 := Person{Name: "dinesh", Item: NewItem(10,20)}

	movers := []Mover{&i, &p1}

	for i, m := range movers {
		fmt.Println(m.Move(i*100,i*200))
	}
}

type Mover interface{
	Move(int , int) string
}

type Person struct {
	Item
	Name string
}

type Item struct {
	X int
	Y int
}

func NewItem(x, y int) Item {
	return Item{
		X: x,
		Y: y,
	}
}

func (i *Item) Move(x, y int) string {
	i.X = x
	i.Y = y

	return "moved to " + string(x) + " and " + string(y)
}

