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
		zookyper.FixCage(cage)
	}
	fmt.Printf("Cage number %d state =>%t\n", cage.Number, cage.IsBroken)

	var sirko = Animal{}.New("Sirko", cage, Cat)
	fmt.Printf("Animal with name %s is sitting in cage number %d type of %s\n", sirko.Name, sirko.Cage.Number, sirko.AnimalType)

	var sirkoSon = Animal{}.New("Sirko junior", cage, Cat)
	fmt.Printf("Animal with name %s is sitting in cage number %d type of %s\n", sirkoSon.Name, sirkoSon.Cage.Number, sirkoSon.AnimalType)

	var sirkoDaughter = Animal{}.New("", nil, Cat)
	if sirkoDaughter == nil {
		fmt.Println("Everything work as expected")
	}

}

const Cat = "Cat"

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
func (z Zookeeper) FixCage(c *Cage) {
	if c.IsBroken {
		c.IsBroken = false
	}
}

type Animal struct {
	Name       string
	Cage       *Cage
	AnimalType string
}

// example nil using
func (a Animal) New(name string, cage *Cage, animalType string) *Animal {
	if name == "" {
		return nil
	}

	return &Animal{
		Name:       name,
		Cage:       cage,
		AnimalType: animalType,
	}
}
