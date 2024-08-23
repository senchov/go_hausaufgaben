package security

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

type Conditions struct {
	IsNight      bool
	DistToAnimal float32
	Battery      int
}

func NewConditions(isNight bool, dst float32, battery int) *Conditions {
	return &Conditions{
		IsNight:      isNight,
		DistToAnimal: dst,
		Battery:      battery,
	}
}

type PTZCamera struct {
	Range float32
	Angle float32
}

func NewPTZCamera(r float32, a float32) *PTZCamera {
	return &PTZCamera{
		Range: r,
		Angle: a,
	}
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

func NewFixedCamera(r float32, a float32) *FixedCamera {
	return &FixedCamera{
		Range: r,
		Angle: a,
	}
}

func (ptz *FixedCamera) Rotate(angle float32) {
	ptz.Angle = 0
}

type NightvisionCamera struct {
	Range float32
	Angle float32
}

func NewNightvisionCamera(r float32, a float32) *NightvisionCamera {
	return &NightvisionCamera{
		Range: r,
		Angle: a,
	}
}

func (nv *NightvisionCamera) Record(c Conditions, records *[]AnimalRecord, rec AnimalRecord) {
	if c.IsNight && nv.Range > c.DistToAnimal && c.Battery > 0 {
		*records = append(*records, rec)
	}
}
