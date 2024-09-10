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

	PrintLocations(rodents, movements)

	now := time.Now()
	movements = moveThroughCenter(1, now.Add(10*time.Minute), WestSector, EastSector, rodents, movements)
	movements = moveThroughCenter(2, now.Add(11*time.Minute), EastSector, NorthSector, rodents, movements)
	movements = moveThroughCenter(2, now.Add(12*time.Minute), NorthSector, SouthSector, rodents, movements)
	movements = moveThroughCenter(3, now.Add(13*time.Minute), SouthSector, WestSector, rodents, movements)

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
	rodents []Rodent,
	movements []Movement,
) []Movement {
	rIndex := Index(rodents, id)
	rodents[rIndex].FromTo[0] = from
	rodents[rIndex].FromTo[1] = to

	movements = append(movements, Movement{MoveTime: moveTime, FromTo: [2]Sector{from, to}})
	return movements
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
