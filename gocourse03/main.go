package main

import (
	"fmt"
	"time"
)

const (
	Chip        = "Chip"
	Dale        = "Dale"
	Cheaser     = "Cheaser"
	Hackwrench  = "Hackwrench"
	NorthSector = "North"
	SouthSector = "South"
	WestSector  = "West"
	EastSector  = "East"
	Center      = "Center"
)

var rodents [4]Rodent
var westRodents []Rodent
var eastRodents []Rodent
var northRodents []Rodent
var southRodents []Rodent

func main() {
	fmt.Println("Start")

	rodents = [4]Rodent{
		{ID: 1, Sector: westRodents, RodentType: Chip, FromTo: [2]string{WestSector, WestSector}},
		{ID: 2, Sector: eastRodents, RodentType: Dale, FromTo: [2]string{EastSector, EastSector}},
		{ID: 3, Sector: southRodents, RodentType: Cheaser, FromTo: [2]string{EastSector, EastSector}},
		{ID: 4, Sector: northRodents, RodentType: Hackwrench, FromTo: [2]string{EastSector, EastSector}},
	}

	westRodents = AddToSecotArr(1, westRodents)
	eastRodents = AddToSecotArr(2, eastRodents)
	westRodents = AddToSecotArr(3, southRodents)
	westRodents = AddToSecotArr(4, northRodents)

	// westRodents = append(westRodents, Rodent{ID: 1, Sector: westRodents, RodentType: Chip, FromTo: [2]string{WestSector, WestSector}})
	// westRodents = RemoveFromSectorArr(1, westRodents)
	// fmt.Printf("%d", len(westRodents))
	//var movements []Movement

}

type Rodent struct {
	ID         int
	Sector     []Rodent
	RodentType string
	FromTo     [2]string
}

type Movement struct {
	Time   time.Time
	FromTo [2]string
}

func moveThroughtCenter(id int, time time.Time, from string, to string) {
	rodent := FindRodent(id)
	rodent.FromTo = [2]string{from, to}
	rodent.Sector = RemoveFromSectorArr(id, rodent.Sector)
	rodent.Sector = GetSectorArr(to)
}

func FindRodent(id int) *Rodent {
	for i := 0; i < len(rodents); i++ {
		if rodents[i].ID == id {
			return &rodents[i]
		}
	}
	return nil
}

func GetSectorArr(name string) []Rodent {
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

func AddToSecotArr(id int, arr []Rodent) []Rodent {
	index := GetIndex(arr, id)
	arr = append(arr, arr[index])
	return arr
}

func RemoveFromSectorArr(id int, arr []Rodent) []Rodent {
	index := GetIndex(arr, id)
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func GetIndex(arr []Rodent, id int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i].ID == id {
			return i
		}
	}
	return -1
}

func PrintLocations() {

}
