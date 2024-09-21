package main

import (
	"fmt"
	"math/rand"
)

const (
	Ungulates AnimalType = "Ungulates"
	Feathered AnimalType = "Feathered"
	Mammals   AnimalType = "Mammals"

	Horses  Breed = "Horses"
	Pony    Breed = "Pony"
	Ducks   Breed = "Ducks"
	Warthog Breed = "Warthog"
	Meerkat Breed = "Meerkat"
)

func main() {
	zoo := NewZoo(CreateAreas())
	timon := zoo.FindAnimal("Timon")
	timon.Eat()
	timon.PrintAnimal()

	zoo.MoveAnimal("Pumbaa", Meerkat)
	animalsInMeerkat := *zoo.Areas[Mammals].Sectors[Meerkat].Animals
	for _, v := range animalsInMeerkat {
		fmt.Println(v)
	}

	for _, v := range zoo.Areas {
		for _, k := range v.Sectors {
			for _, a := range *k.Animals {
				isFeed := rand.Intn(2) == 1
				if isFeed {
					a.Eat()
				}
				a.PrintAnimal()
			}
		}
	}
}

func (a Animal) PrintAnimal() {
	var hunger string
	if a.IsAte {
		hunger = "in not hunger"
	} else {
		hunger = "is hunger"
	}
	fmt.Printf("Animal with name %s %s\n", a.Name, hunger)
}

type Areas map[AnimalType]Area

type Area struct {
	Number  int
	Sectors map[Breed]Sector
}

type Sector struct {
	Subtype Breed
	Animals *[]Animal
}

type Animal struct {
	ID    int
	Name  string
	IsAte bool
}

type Zoo struct {
	Areas Areas
}

type AnimalType string
type Breed string

func NewZoo(areas map[AnimalType]Area) *Zoo {
	return &Zoo{Areas: areas}
}

func CreateAreas() map[AnimalType]Area {
	areas := map[AnimalType]Area{
		Ungulates: Area{
			Number: 1,
			Sectors: map[Breed]Sector{
				Horses: Sector{
					Subtype: Horses,
					Animals: &[]Animal{
						{ID: 1, Name: "Mike"},
						{ID: 2, Name: "Roach"},
					},
				},
				Pony: Sector{
					Subtype: Pony,
					Animals: &[]Animal{
						{ID: 3, Name: "Arnold"},
					},
				},
			},
		},
		Feathered: Area{
			Number: 2,
			Sectors: map[Breed]Sector{
				Ducks: Sector{
					Subtype: Ducks,
					Animals: &[]Animal{
						{ID: 4, Name: "Donald"},
					},
				},
			},
		},
		Mammals: Area{
			Number: 3,
			Sectors: map[Breed]Sector{
				Warthog: Sector{
					Subtype: Warthog,
					Animals: &[]Animal{
						{ID: 5, Name: "Pumbaa"},
					},
				},
				Meerkat: Sector{
					Subtype: Meerkat,
					Animals: &[]Animal{
						{ID: 6, Name: "Timon"},
					},
				},
			},
		},
	}
	return areas
}

func (zoo *Zoo) FindAnimal(name string) *Animal {
	for _, area := range zoo.Areas {
		for _, sector := range area.Sectors {
			for _, a := range *sector.Animals {
				if a.Name == name {
					return &a
				}
			}
		}
	}

	return nil
}

func (zoo *Zoo) FindAnimalSector(name string) *Sector {
	for _, area := range zoo.Areas {
		for _, sector := range area.Sectors {
			for _, a := range *sector.Animals {
				if a.Name == name {
					return &sector
				}
			}
		}
	}

	return nil
}

func (zoo *Zoo) FindSector(sector Breed) *Sector {
	for _, area := range zoo.Areas {
		v, ok := area.Sectors[sector]
		if ok {
			return &v
		}
	}
	return nil
}

func AnimalIndex(animals []Animal, id int) int {
	for i := 0; i < len(animals); i++ {
		if animals[i].ID == id {
			return i
		}
	}
	return -1
}

func (zoo *Zoo) MoveAnimal(name string, sector Breed) {
	a := zoo.FindAnimal(name)
	if a == nil {
		fmt.Errorf("Animal with name %s is absent in zoo", name)
		return
	}
	sTarget := zoo.FindSector(sector)
	if sTarget == nil {
		fmt.Errorf("Sector with breed %s is absent", sector)
		return
	}
	sTarget.AddAnimal(*a)

	sSource := zoo.FindAnimalSector(name)
	if sSource == nil {
		fmt.Errorf("Sector for animal %s is absent", name)
		return
	}

	sourсeIndex := AnimalIndex(*sSource.Animals, a.ID)
	if sourсeIndex == -1 {
		fmt.Errorf("Animal with id %d is absent", a.ID)
	}

	*sSource.Animals = append((*sSource.Animals)[:sourсeIndex], (*sSource.Animals)[sourсeIndex+1:]...)
}

func (sector *Sector) AddAnimal(animal Animal) {
	*sector.Animals = append(*sector.Animals, animal)
}

func (a *Animal) Eat() {
	a.IsAte = true
}
