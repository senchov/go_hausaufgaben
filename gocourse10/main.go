package main

import (
	"fmt"

	"github.com/senchov/go_hausaufgaben/gocourse10/observer"
)

func main() {
	fmt.Println("Start")
	// strategy
	collar := Collar{Trough: MeetTrought{}, ID: 1}
	Feed(collar)
	collar.SetTrough(&FishTrough{})
	Feed(collar)
	collar.SetTrough(&VegetablesTrought{})
	Feed(collar)

	//decorator
	collar.SetTrough(&SousegesTrough{Trough: MeetTrought{}})
	Feed(collar)

	item := observer.Item{}
	item.Register(collar)

	for i := 0; i < 24; i++ {
		if i == 9 {
			fmt.Println("Go to breakfast")
			item.NotifyAll()
			collar.SetTrough(&VegetablesTrought{})
		}

		if i == 13 {
			fmt.Println("Go to the lunch")
			item.NotifyAll()
			collar.SetTrough(&FishTrough{})
		}

		if i == 18 {
			fmt.Println("Go to the dinner")
			item.NotifyAll()
			collar.SetTrough(&MeetTrought{})
		}
	}
}

func Feed(c Collar) {
	fmt.Printf("Feed with %s \n", c.Trough.Food())
}

type Trough interface {
	Food() string
}

type MeetTrought struct {
}

func (MeetTrought) Food() string {
	return "Ham, pork and chicken"
}

type VegetablesTrought struct {
}

func (VegetablesTrought) Food() string {
	return "Carrots, potatoes and cucumbers"
}

type FishTrough struct {
}

func (FishTrough) Food() string {
	return "Code and salmon"
}

type Collar struct {
	Trough
	ID int
}

func (c *Collar) SetTrough(t Trough) {
	c.Trough = t
}

type SousegesTrough struct {
	Trough
}

func (st SousegesTrough) Food() string {
	return "Sauseges from " + st.Trough.Food()
}

func (c Collar) TimeToEat() {
	Feed(c)
}

func (c Collar) GetID() int {
	return c.ID
}
