package main

import (
	"fmt"
)

func main() {
	i := NewItem(10, 20)
	p1 := Player{Name: "dinesh", Item: NewItem(10, 20)}

	movers := []Mover{&i, &p1}
	for _, v := range movers {
		switch v.(type) {
		case *Item:
			fmt.Println("Item type", v.(*Item).X)
		case *Player:
			fmt.Println("Player type", v.(*Player).Keys)
		default:
			fmt.Printf("unknown %T\n", v)
		}
	}
}

type Mover interface {
	Move(int, int)
}

type Player struct {
	Item
	Name string
	Keys []Key
}

type Key byte

const (
	Bangaram = iota + 1 //Gold
	Vendi               //Silver
	Ithhadi             //Bronze
	Raagi               //Raagi
	invalidKey
)

func (k Key) String() string {
	switch k {
	case Bangaram:
		return "Gold"
	case Vendi:
		return "Silver"
	case Ithhadi:
		return "Bronze"
	case Raagi:
		return "Raagi"
	}
	return fmt.Sprintf("invalid key %d", k)
}

func (p *Player) FoundKey(k Key) error {
	if k < Bangaram || k >= invalidKey {
		return fmt.Errorf("invalid key %d", k)
	}

	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}

	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, k1 := range keys {
		if k1 == k {
			return true
		}
	}
	return false
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

func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}
