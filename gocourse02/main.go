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

	var sirko = Animal{}.New("Sirko", cage)
	fmt.Printf("Animal with name %s is istting in cage number %d\n", sirko.Name, sirko.Cage.Number)

	var sirkoSon = Animal{}.New("Sirko junior", cage)
	fmt.Printf("Animal with name %s is istting in cage number %d\n", sirkoSon.Name, sirkoSon.Cage.Number)

	var sirkoDaughter = Animal{}.New("", nil)
	if sirkoDaughter == nil {
		fmt.Println("Everything work as expected")
	}

}

// example with struct
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

// example with constructor
func (c Cage) New(isBroken bool, number int) *Cage {
	return &Cage{
		IsBroken: isBroken,
		Number:   number,
	}
}

// example with pointer
func (c *Cage) FixCage() {
	if c.IsBroken {
		c.IsBroken = false
	}
}

type Animal struct {
	Name string
	Cage *Cage
}

// example nil using
func (a Animal) New(name string, cage *Cage) *Animal {
	if name == "" {
		return nil
	}

	return &Animal{
		Name: name,
		Cage: cage,
	}
}
