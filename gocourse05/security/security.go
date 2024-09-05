package security

import "fmt"

type Range float32
type Angle float32

type Rotater interface {
	Rotate(angle Angle)
}

type Recorder interface {
	Record(c Condition, records *[]AnimalRecord, rec AnimalRecord)
}

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

type Condition struct {
	IsNight      bool
	DistToAnimal Range
	Battery      int
}

func NewConditions(isNight bool, dst Range, battery int) *Condition {
	return &Condition{
		IsNight:      isNight,
		DistToAnimal: dst,
		Battery:      battery,
	}
}

type PTZCamera struct {
	Range
	Angle
}

func NewPTZCamera(r Range, a Angle) *PTZCamera {
	return &PTZCamera{
		Range: r,
		Angle: a,
	}
}

func (ptz *PTZCamera) Rotate(angle Angle) {
	ptz.Angle += angle
}

func (ptz *PTZCamera) Record(c Condition, records *[]AnimalRecord, rec AnimalRecord) {
	if !c.IsNight && ptz.Range > c.DistToAnimal && c.Battery > 0 {
		*records = append(*records, rec)
	}
}

type FixedCamera struct {
	Range
	Angle
}

func NewFixedCamera(r Range, a Angle) *FixedCamera {
	return &FixedCamera{
		Range: r,
		Angle: a,
	}
}

func (ptz *FixedCamera) Rotate(angle float32) {
	ptz.Angle = 0
}

type NightvisionCamera struct {
	Range
	Angle
}

func NewNightvisionCamera(r Range, a Angle) *NightvisionCamera {
	return &NightvisionCamera{
		Range: r,
		Angle: a,
	}
}

func (nv *NightvisionCamera) Record(c Condition, records *[]AnimalRecord, rec AnimalRecord) {
	if c.IsNight && nv.Range > c.DistToAnimal && c.Battery > 0 {
		*records = append(*records, rec)
	}
}

func (nv *NightvisionCamera) Rotate(angle Angle) {
	nv.Angle += angle
}
