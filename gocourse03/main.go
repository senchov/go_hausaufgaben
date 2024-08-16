package main

import (
	"fmt"
	"time"
)

const (
	Chip        RodentType = "Chip"
	Dale        RodentType = "Dale"
	Cheaser     RodentType = "Cheaser"
	Hackwrench  RodentType = "Hackwrench"
	NorthSector Sector     = "North"
	SouthSector Sector     = "South"
	WestSector  Sector     = "West"
	EastSector  Sector     = "East"
	Center      Sector     = "Center"
)

type RodentType string

type Sector string

func main() {
	var rodents [4]Rodent
	var westRodents []Rodent
	var eastRodents []Rodent
	var northRodents []Rodent
	var southRodents []Rodent
	var movements []Movement

	rodents = [4]Rodent{
		{ID: 1, Sector: &westRodents, RodentType: Chip, FromTo: [2]Sector{WestSector, WestSector}},
		{ID: 2, Sector: &eastRodents, RodentType: Dale, FromTo: [2]Sector{EastSector, EastSector}},
		{ID: 3, Sector: &southRodents, RodentType: Cheaser, FromTo: [2]Sector{SouthSector, SouthSector}},
		{ID: 4, Sector: &northRodents, RodentType: Hackwrench, FromTo: [2]Sector{NorthSector, NorthSector}},
	}

	westRodents = append(westRodents, rodents[0])
	eastRodents = append(eastRodents, rodents[1])
	southRodents = append(southRodents, rodents[2])
	northRodents = append(northRodents, rodents[3])

	PrintLocations(rodents, movements)

	moveThroughtCenter(1, time.Now().Add(10), WestSector, EastSector, movements, rodents, westRodents, eastRodents, northRodents, southRodents)
	moveThroughtCenter(2, time.Now().Add(11), EastSector, NorthSector, movements, rodents, westRodents, eastRodents, northRodents, southRodents)
	moveThroughtCenter(2, time.Now().Add(12), NorthSector, SouthSector, movements, rodents, westRodents, eastRodents, northRodents, southRodents)
	moveThroughtCenter(3, time.Now().Add(13), SouthSector, WestSector, movements, rodents, westRodents, eastRodents, northRodents, southRodents)

	PrintLocations(rodents, movements)
}

type Rodent struct {
	ID         int
	Sector     *[]Rodent
	RodentType RodentType
	FromTo     [2]Sector
}

type Movement struct {
	MoveTime time.Time
	FromTo   [2]Sector
}

func moveThroughtCenter(
	id int,
	moveTime time.Time,
	from Sector,
	to Sector,
	movements []Movement,
	rodents [4]Rodent,
	westRodents []Rodent,
	eastRodents []Rodent,
	northRodents []Rodent,
	southRodents []Rodent,
) {
	rodent := FindRodent(id, rodents)
	rodent.FromTo = [2]Sector{from, to}

	rodent.Sector = RemoveFromSectorArr(id, *rodent.Sector)
	rodent.Sector = GetSectorArr(to, &westRodents, &eastRodents, &northRodents, &southRodents)

	var c []Rodent = *rodent.Sector
	c = append(c, *rodent)
	rodent.Sector = &c
	movements = append(movements, Movement{MoveTime: moveTime, FromTo: [2]Sector{WestSector, EastSector}})
}

func FindRodent(id int, rodents [4]Rodent) *Rodent {
	for i := 0; i < len(rodents); i++ {
		if rodents[i].ID == id {
			return &rodents[i]
		}
	}
	return nil
}

func GetSectorArr(
	name Sector,
	westRodents *[]Rodent,
	eastRodents *[]Rodent,
	northRodents *[]Rodent,
	southRodents *[]Rodent) *[]Rodent {
	switch name {
	case WestSector:
		return westRodents
	case EastSector:
		return eastRodents
	case NorthSector:
		return northRodents
	case SouthSector:
		return southRodents
	default:
		return nil
	}
}

func RemoveFromSectorArr(id int, arr []Rodent) *[]Rodent {
	index := GetIndex(arr, id)
	arr = append(arr[:index], arr[index+1:]...)
	return &arr
}

func GetIndex(arr []Rodent, id int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i].ID == id {
			return i
		}
	}
	return -1
}

func PrintLocations(rodents [4]Rodent, movements []Movement) {
	for _, val := range rodents {
		fmt.Printf("%s is located in %s sector \n", val.RodentType, val.FromTo[1])
	}

	for _, val := range movements {
		fmt.Printf("At time %v move from %s to %s \n", val.MoveTime, val.FromTo[0], val.FromTo[1])
	}
	fmt.Println("----------------")
}
