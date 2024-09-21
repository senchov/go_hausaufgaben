package observer

import "slices"

type Observer interface {
	TimeToEat()
	GetID() int
}

type Item struct {
	observers []Observer
}

func (i *Item) Register(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *Item) Deregister(o Observer) {
	i.observers = slices.DeleteFunc(i.observers, func(e Observer) bool {
		return e.GetID() == o.GetID()
	})
}

func (i *Item) NotifyAll() {
	for _, observer := range i.observers {
		observer.TimeToEat()
	}
}
