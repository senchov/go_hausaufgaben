package main

import (
	"fmt"
)

func main() {
	zookyper := Zookyper{
		Name:   "Mukola",
		IsBusy: false,
	}
	zookyper.PrintName()

	cage := Cage{
		Number:   1,
		IsBroken: true,
	}

	fmt.Printf("Cage number %d state =>%t\n", cage.Number, cage.IsBroken)
	if !zookyper.IsBusy && cage.IsBroken {
		cage.FixCage()
	}
	fmt.Printf("Cage number %d state =>%t\n", cage.Number, cage.IsBroken)

	var sirko = Animal{
		Name:   "Sirko",
		Number: 1,
	}
	fmt.Printf("Animal with name %s is istting in cage number %d\n", sirko.Name, sirko.Number)

	var sirko_son = sirko.Duplicate("Sirko junior", 1)
	fmt.Printf("Animal with name %s is istting in cage number %d\n", sirko_son.Name, sirko_son.Number)

}

type Zookyper struct {
	IsBusy bool
	Name   string
}

func (z Zookyper) PrintName() {
	fmt.Printf("Zookyper name is %v \n", z.Name)
}

type Cage struct {
	IsBroken bool
	Number   int32
}

func (c *Cage) FixCage() {
	if c.IsBroken {
		c.IsBroken = false
	}

}

type Animal struct {
	Name   string
	Number int32
}

func (a Animal) Duplicate(name string, number int32) Animal {
	return Animal{
		Name:   name,
		Number: number,
	}
}
