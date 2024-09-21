package feed

import (
	"log"
)

func Feed(c Collar) {
	log.Printf("Feed with %s \n", c.Trough.Food())
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

type SausagesTrough struct {
	Trough
}

func (st SausagesTrough) Food() string {
	return "Sausages from " + st.Trough.Food()
}

func (c Collar) TimeToEat() {
	Feed(c)
}

func (c Collar) GetID() int {
	return c.ID
}
