package animal

import "fmt"

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

type AnimalRecord struct {
	AnimalType
	AnimalAction
}

func (r AnimalRecord) Print() {
	fmt.Printf("Animal %s is doing %s \n", r.AnimalType, r.AnimalAction)
}

func NewAnimalRecord(t AnimalType, a AnimalAction) *AnimalRecord {
	return &AnimalRecord{
		AnimalType:   t,
		AnimalAction: a,
	}
}
