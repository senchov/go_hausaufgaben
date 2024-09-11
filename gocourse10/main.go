package main

import "fmt"

func main() {
	fmt.Println("Start")
	// strategy
	collar := Collar{Trough: MeetTrought{}}
	Feed(collar)
	collar.SetTrough(&FishTrough{})
	Feed(collar)
	collar.SetTrough(&VegetablesTrought{})
	Feed(collar)

	//decorator
	collar.SetTrough(&SousegesTrough{Trough: MeetTrought{}})
	Feed(collar)
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
