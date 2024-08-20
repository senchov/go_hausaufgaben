package main

import "fmt"

func main() {
	records := []AnimalRecord{}

	ptz := PTZCamera{Angle: 0, Range: 10}
	nv := NightvisionCamera{Angle: 30, Range: 6}

	ptz.Record(Conditions{IsNight: false, DistToAnimal: 2, Battery: 100}, &records, AnimalRecord{Bear, Sleep})
	ptz.Record(Conditions{IsNight: true, DistToAnimal: 4, Battery: 100}, &records, AnimalRecord{Bear, Sleep})
	nv.Record(Conditions{IsNight: true, DistToAnimal: 5, Battery: 100}, &records, AnimalRecord{Tiger, Walking})
	ptz.Record(Conditions{IsNight: false, DistToAnimal: 5, Battery: 100}, &records, AnimalRecord{Rat, Eat})
	nv.Record(Conditions{IsNight: true, DistToAnimal: 50, Battery: 100}, &records, AnimalRecord{Tiger, Walking})
	nv.Record(Conditions{IsNight: true, DistToAnimal: 3, Battery: 100}, &records, AnimalRecord{Tiger, Walking})

	for _, v := range records {
		fmt.Println(v)
	}
}

type Rotater interface {
	Rotate(angle float32)
}

type Recorder interface {
	Record(c Conditions, records *[]AnimalRecord, rec AnimalRecord)
}

type AnimalRecord struct {
	AnimalType
	AnimalAction
}

type AnimalType string
type AnimalAction string

const (
	Bear  AnimalType = "Bear"
	Tiger AnimalType = "Tiger"
	Rat   AnimalType = "Rat"

	Sleep   AnimalAction = "Sleep"
	Eat     AnimalAction = "Eat"
	Walking AnimalAction = "Walking"
)

type Conditions struct {
	IsNight      bool
	DistToAnimal float32
	Battery      int
}

type PTZCamera struct {
	Range float32
	Angle float32
}

func (ptz *PTZCamera) Rotate(angle float32) {
	ptz.Angle += angle
}

func (ptz *PTZCamera) Record(c Conditions, records *[]AnimalRecord, rec AnimalRecord) {
	if !c.IsNight && ptz.Range > c.DistToAnimal && c.Battery > 0 {
		*records = append(*records, rec)
	}
}

type FixedCamera struct {
	Range float32
	Angle float32
}

func (ptz *FixedCamera) Rotate(angle float32) {
	ptz.Angle = 0
}

type NightvisionCamera struct {
	Range float32
	Angle float32
}

func (nv *NightvisionCamera) Record(c Conditions, records *[]AnimalRecord, rec AnimalRecord) {
	if c.IsNight && nv.Range > c.DistToAnimal && c.Battery > 0 {
		*records = append(*records, rec)
	}
}
