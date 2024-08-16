package main

import (
	"fmt"
)

func main() {
	zookeeper := NewZookeeper(false, "Mukola")
	zookeeper.PrintName()

	cage := NewCage(true, 1)
	fmt.Printf("Cage number %d state =>%t\n", cage.Number, cage.IsBroken)
	if !zookeeper.IsBusy && cage.IsBroken {
		zookeeper.FixCage(cage)
	}
	fmt.Printf("Cage number %d state =>%t\n", cage.Number, cage.IsBroken)

	var sirko = NewAnimal("Sirko", cage, Cat)
	fmt.Printf("Animal with name %s is sitting in cage number %d type of %s\n", sirko.Name, sirko.Cage.Number, sirko.AnimalType)

	var sirkoSon = NewAnimal("Sirko junior", cage, Tiger)
	fmt.Printf("Animal with name %s is sitting in cage number %d type of %s\n", sirkoSon.Name, sirkoSon.Cage.Number, sirkoSon.AnimalType)

	var sirkoDaughter = NewAnimal("", nil, Manul)
	if sirkoDaughter == nil {
		fmt.Println("Everything work as expected")
	}

}

const Cat = "Cat"
const Manul = "Manul"
const Tiger = "Tiger"

type Zookeeper struct {
	IsBusy bool
	Name   string
}

func NewZookeeper(isBusy bool, name string) *Zookeeper {
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

func NewCage(isBroken bool, number int) *Cage {
	return &Cage{
		IsBroken: isBroken,
		Number:   number,
	}
}

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

func NewAnimal(name string, cage *Cage, animalType string) *Animal {
	if name == "" {
		return nil
	}

	return &Animal{
		Name:       name,
		Cage:       cage,
		AnimalType: animalType,
	}
}
