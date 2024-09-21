package aquarium

type Aquarium struct {
	Size int
	Animal
	SaltLevel int
	Chemicals []Chemicals
}

type Animal string
type Chemicals string

const (
	Cat   Animal = "Cat"
	Manul Animal = "Manul"
	Tiger Animal = "Tiger"

	N  Chemicals = "N"
	Mg Chemicals = "Mg"
	Na Chemicals = "Na"
	K  Chemicals = "K"
)

func New(options ...func(*Aquarium)) *Aquarium {
	aq := &Aquarium{}
	for _, v := range options {
		v(aq)
	}
	return aq
}

func WithSize(size int) func(*Aquarium) {
	return func(a *Aquarium) {
		a.Size = size
	}
}

func WithAnimal(animal Animal) func(*Aquarium) {
	return func(a *Aquarium) {
		a.Animal = animal
	}
}

func WithSaltLevel(saltLevel int) func(*Aquarium) {
	return func(a *Aquarium) {
		a.SaltLevel = saltLevel
	}
}

func WithChemicals(c []Chemicals) func(*Aquarium) {
	return func(a *Aquarium) {
		a.Chemicals = c
	}
}
