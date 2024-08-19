package main

import "fmt"

// Наприклад, зоопарк ділиться на території: копитні, пернаті, примати… та інше.

const (
	Ungulates AnimalType = "Ungulates"
	Feathered AnimalType = "Feathered"
	Mammals   AnimalType = "Mammals"

	Horses   Breed = "Horses"
	Pony     Breed = "Pony"
	Ducks    Breed = "Ducks"
	Warthoag Breed = "Warthog"
	Meerkat  Breed = "Meerkat"
)

func main() {
	fmt.Println("Start")
	zoo := NewZoo()
	fmt.Println(zoo)
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

func NewZoo() *Zoo {
	zoo := Zoo{
		Areas: map[AnimalType]Area{
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
					Warthoag: Sector{
						Subtype: Warthoag,
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
		},
	}

	return &zoo
}

// Реалізувати функції пошуку тварини за імʼям або ID, переміщення тварини із загону в загін, годування тварини.
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

func AnimalIndex(slice []Animal, id int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i].ID == id {
			return i
		}
	}
	return -1
}

func (zoo *Zoo) MoveAnimal(name string, sector Breed) {
	a := zoo.FindAnimal(name)
	sTarget := zoo.FindSector(sector)
	AddAnimalToSector(sTarget, a)

	sSource := zoo.FindAnimalSector(name)
	sourseIndex := AnimalIndex(*sSource.Animals, a.ID)
	*sSource.Animals = append((*sSource.Animals)[:sourseIndex], (*sSource.Animals)[sourseIndex+1:]...)
}

func AddAnimalToSector(sector *Sector, animal *Animal) {
	*sector.Animals = append(*sector.Animals, *animal)
}

func (a *Animal) FeedAnimal() {
	a.IsAte = true
}
