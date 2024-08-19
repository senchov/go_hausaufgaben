package main

import (
	"fmt"
	"time"
)

const (
	Chip       RodentType = "Chip"
	Dale       RodentType = "Dale"
	Cheaser    RodentType = "Cheaser"
	Hackwrench RodentType = "Hackwrench"

	NorthSector Sector = "North"
	SouthSector Sector = "South"
	WestSector  Sector = "West"
	EastSector  Sector = "East"
	Center      Sector = "Center"
)

type RodentType string

type Sector string

func main() {
	var movements []Movement

	rodents := [4]Rodent{
		{ID: 1, RodentType: Chip, FromTo: [2]Sector{WestSector, WestSector}},
		{ID: 2, RodentType: Dale, FromTo: [2]Sector{EastSector, EastSector}},
		{ID: 3, RodentType: Cheaser, FromTo: [2]Sector{SouthSector, SouthSector}},
		{ID: 4, RodentType: Hackwrench, FromTo: [2]Sector{NorthSector, NorthSector}},
	}

	westRodents := []Rodent{rodents[0]}
	eastRodents := []Rodent{rodents[1]}
	southRodents := []Rodent{rodents[2]}
	northRodents := []Rodent{rodents[3]}

	PrintLocations(&rodents, &movements)

	moveThroughCenter(1, time.Now().Add(10), WestSector, EastSector, &movements, &rodents, &westRodents, &eastRodents, &northRodents, &southRodents)
	moveThroughCenter(2, time.Now().Add(11), EastSector, NorthSector, &movements, &rodents, &westRodents, &eastRodents, &northRodents, &southRodents)
	moveThroughCenter(2, time.Now().Add(12), NorthSector, SouthSector, &movements, &rodents, &westRodents, &eastRodents, &northRodents, &southRodents)
	moveThroughCenter(3, time.Now().Add(13), SouthSector, WestSector, &movements, &rodents, &westRodents, &eastRodents, &northRodents, &southRodents)

	PrintLocations(&rodents, &movements)
}

type Rodent struct {
	ID         int
	RodentType RodentType
	FromTo     [2]Sector
}

type Movement struct {
	MoveTime time.Time
	FromTo   [2]Sector
}

func moveThroughCenter(
	id int,
	moveTime time.Time,
	from Sector,
	to Sector,
	movements *[]Movement,
	rodents *[4]Rodent,
	westRodents *[]Rodent,
	eastRodents *[]Rodent,
	northRodents *[]Rodent,
	southRodents *[]Rodent,
) {
	rIndex := Index(rodents[:], id)
	rodents[rIndex].FromTo[0] = from
	rodents[rIndex].FromTo[1] = to

	dst := SectorSlice(to, westRodents, eastRodents, northRodents, southRodents)
	*dst = append(*dst, rodents[rIndex])

	src := SectorSlice(from, westRodents, eastRodents, northRodents, southRodents)
	index := Index(*src, id)
	*src = append((*src)[:index], (*src)[index+1:]...)

	*movements = append(*movements, Movement{MoveTime: moveTime, FromTo: [2]Sector{WestSector, EastSector}})
}

func SectorSlice(
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

func Index(slice []Rodent, id int) int {
	for i, v := range slice {
		if v.ID == id {
			return i
		}
	}

	return -1
}

func PrintLocations(rodents *[4]Rodent, movements *[]Movement) {
	for _, val := range rodents {
		fmt.Printf("%s is located in %s sector \n", val.RodentType, val.FromTo[1])
	}

	for _, val := range *movements {
		fmt.Printf("At time %v move from %s to %s \n", val.MoveTime, val.FromTo[0], val.FromTo[1])
	}
	fmt.Println("----------------")
}
