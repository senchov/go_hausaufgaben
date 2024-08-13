package main

import (
	"fmt"
)

func main() {
	zookyper := Zookeeper{}.New(false, "Mukola")
	zookyper.PrintName()

	cage := Cage{}.New(true, 1)
	fmt.Printf("Cage number %d state =>%t\n", cage.Number, cage.IsBroken)
	if !zookyper.IsBusy && cage.IsBroken {
		cage.FixCage()
	}
	fmt.Printf("Cage number %d state =>%t\n", cage.Number, cage.IsBroken)

	var sirko = Animal{}.New("Sirko", 1)
	fmt.Printf("Animal with name %s is istting in cage number %d\n", sirko.Name, sirko.Number)

	var sirko_son = Animal{}.New("Sirko junior", 1)
	fmt.Printf("Animal with name %s is istting in cage number %d\n", sirko_son.Name, sirko_son.Number)
}

type Zookeeper struct {
	IsBusy bool
	Name   string
}

func (z Zookeeper) New(isBusy bool, name string) *Zookeeper {
	return &Zookeeper{
		IsBusy: isBusy,
		Name:   name,
	}
}

func (z Zookeeper) PrintName() {
	fmt.Printf("Zookyper name is %v \n", z.Name)
}

type Cage struct {
	IsBroken bool
	Number   int
}

func (c Cage) New(isBroken bool, number int) *Cage {
	return &Cage{
		IsBroken: isBroken,
		Number:   number,
	}
}

func (c *Cage) FixCage() {
	if c.IsBroken {
		c.IsBroken = false
	}
}

type Animal struct {
	Name   string
	Number int
}

func (a Animal) New(name string, number int) *Animal {
	return &Animal{
		Name:   name,
		Number: number,
	}
}
