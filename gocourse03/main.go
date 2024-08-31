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

	NorthSector  Sector = "North"
	SouthSector  Sector = "South"
	WestSector   Sector = "West"
	EastSector   Sector = "East"
	CenterSector Sector = "Center"
)

type RodentType string

type Sector string

func main() {
	var movements []Movement

	rodents := []Rodent{
		{ID: 1, RodentType: Chip, FromTo: [2]Sector{WestSector, WestSector}},
		{ID: 2, RodentType: Dale, FromTo: [2]Sector{EastSector, EastSector}},
		{ID: 3, RodentType: Cheaser, FromTo: [2]Sector{SouthSector, SouthSector}},
		{ID: 4, RodentType: Hackwrench, FromTo: [2]Sector{NorthSector, NorthSector}},
	}

	westRodents := []Rodent{rodents[0]}
	eastRodents := []Rodent{rodents[1]}
	southRodents := []Rodent{rodents[2]}
	northRodents := []Rodent{rodents[3]}

	param := RodentParameter{
		Rodents:      &rodents,
		WestRodents:  &westRodents,
		EastRodents:  &eastRodents,
		SouthRodents: &southRodents,
		NorthRodents: &northRodents,
	}

	PrintLocations(rodents, movements)

	now := time.Now()
	movements = moveThroughCenter(1, now.Add(10*time.Minute), WestSector, EastSector, param, movements)
	movements = moveThroughCenter(2, now.Add(11*time.Minute), EastSector, NorthSector, param, movements)
	movements = moveThroughCenter(2, now.Add(12*time.Minute), NorthSector, SouthSector, param, movements)
	movements = moveThroughCenter(3, now.Add(13*time.Minute), SouthSector, WestSector, param, movements)

	PrintLocations(rodents, movements)
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
	param RodentParameter,
	movements []Movement,
) []Movement {
	rIndex := Index(*param.Rodents, id)
	(*param.Rodents)[rIndex].FromTo[0] = from
	(*param.Rodents)[rIndex].FromTo[1] = to

	dst := FindSectorSlice(to, param.WestRodents, param.EastRodents, param.NorthRodents, param.SouthRodents)
	*dst = append(*dst, (*param.Rodents)[rIndex])

	src := FindSectorSlice(from, param.WestRodents, param.EastRodents, param.NorthRodents, param.SouthRodents)
	index := Index(*src, id)
	*src = append((*src)[:index], (*src)[index+1:]...)

	movements = append(movements, Movement{MoveTime: moveTime, FromTo: [2]Sector{from, to}})
	return movements
}

func FindSectorSlice(
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

func Index(rodents []Rodent, id int) int {
	for i, v := range rodents {
		if v.ID == id {
			return i
		}
	}

	return -1
}

func PrintLocations(rodents []Rodent, movements []Movement) {
	for _, val := range rodents {
		fmt.Printf("%s is located in %s sector \n", val.RodentType, val.FromTo[1])
	}

	for _, val := range movements {
		fmt.Printf("At time %v move from %s to %s \n", val.MoveTime, val.FromTo[0], val.FromTo[1])
	}
	fmt.Println("----------------")
}

type RodentParameter struct {
	Rodents      *[]Rodent
	WestRodents  *[]Rodent
	EastRodents  *[]Rodent
	NorthRodents *[]Rodent
	SouthRodents *[]Rodent
}
