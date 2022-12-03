package rucksack

import (
	"fmt"
	"strings"
)

type Rucksack struct {
	Number         int
	CompartmentOne *Compartment
	CompartmentTwo *Compartment
}

type Compartment struct {
	Items []*Item
}

type Item struct {
	Value string
}

const (
	priorityRange = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Priority int

func (i *Item) Priority() Priority {
	return Priority(strings.Index(priorityRange, i.Value) + 1)
}

func New(number int) *Rucksack {
	rucksack := Rucksack{
		Number:         number,
		CompartmentOne: NewCompartment(),
		CompartmentTwo: NewCompartment(),
	}
	return &rucksack
}

func NewCompartment() *Compartment {
	compartment := Compartment{}
	compartment.Items = make([]*Item, 0)
	return &Compartment{}
}

func NewItem(value string) *Item {
	item := Item{Value: value}
	return &item
}

func (c *Compartment) AddItem(item *Item) {
	//if c.Items == nil {
	//	c.Items = &Items{}
	//}
	c.Items = append(c.Items, item)
}

func (i *Item) String() string {
	return i.Value
}

func (c *Compartment) String() string {
	s := make([]string, 0)
	for _, v := range c.Items {
		s = append(s, v.Value)
	}
	return strings.Join(s, "")
}

func (r *Rucksack) String() string {
	return fmt.Sprintf("Rucksack id :%d\n Compartment one: %s\n Compartment two: %s\n", r.Number,
		r.CompartmentOne.String(), r.CompartmentTwo.String())
}

func (r *Rucksack) GetAllItems() []*Item {
	items := make([]*Item, 0)

	for _, i := range r.CompartmentOne.Items {
		items = append(items, i)
	}

	for _, i := range r.CompartmentTwo.Items {
		items = append(items, i)
	}

	return items
}

// GetDuplicateItems
// Returns the similarities
// /*
func (r *Rucksack) GetDuplicateItems() map[string]*Item {
	duplicates := make(map[string]*Item, 0)
	for _, v1 := range r.CompartmentOne.Items {
		for _, v2 := range r.CompartmentTwo.Items {
			if v1.Value == v2.Value {
				duplicates[v1.Value] = v1
			}
		}
	}

	return duplicates
}
