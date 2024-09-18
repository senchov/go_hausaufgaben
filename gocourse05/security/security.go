package security

import (
	"gocourse05/animal"
)

type Range float32
type Angle float32

type Condition struct {
	IsNight      bool
	DistToAnimal Range
	Battery      int
}

func NewCondition(isNight bool, distToAnimal Range, battery int) *Condition {
	return &Condition{
		IsNight:      isNight,
		DistToAnimal: distToAnimal,
		Battery:      battery,
	}
}

type PTZCamera struct {
	Range
	Angle
}

func NewPTZCamera(cameraRange Range, angle Angle) *PTZCamera {
	return &PTZCamera{
		Range: cameraRange,
		Angle: angle,
	}
}

func (ptz *PTZCamera) Rotate(angle Angle) {
	ptz.Angle += angle
}

func (ptz *PTZCamera) Record(condition Condition, records []animal.AnimalRecord, record animal.AnimalRecord) []animal.AnimalRecord {
	if !condition.IsNight && ptz.Range > condition.DistToAnimal && condition.Battery > 0 {
		records = append(records, record)
	}
	return records
}

type FixedCamera struct {
	Range
}

func NewFixedCamera(cameraRange Range) *FixedCamera {
	return &FixedCamera{
		Range: cameraRange,
	}
}

func (fc *FixedCamera) Record(condition Condition, records []animal.AnimalRecord, record animal.AnimalRecord) []animal.AnimalRecord {
	if condition.IsNight && fc.Range > condition.DistToAnimal && condition.Battery > 0 {
		records = append(records, record)
	}
	return records
}

type NightvisionCamera struct {
	Range
	Angle
}

func NewNightvisionCamera(cameraRange Range, angele Angle) *NightvisionCamera {
	return &NightvisionCamera{
		Range: cameraRange,
		Angle: angele,
	}
}

func (nv *NightvisionCamera) Record(condition Condition, records []animal.AnimalRecord, record animal.AnimalRecord) []animal.AnimalRecord {
	if condition.IsNight && nv.Range > condition.DistToAnimal && condition.Battery > 0 {
		records = append(records, record)
	}
	return records
}

func (nv *NightvisionCamera) Rotate(angle Angle) {
	nv.Angle += angle
}
