package aquarium

// У тераріумі зоопарку купа акваріумів, і всі різні: розмір, тварина всередині, рівень солі. Кожен акваріум віддає на центральний сервер хімічний склад води.
type Aquarium struct {
	Size int
	Animal
	SaltLevel  int
	Chemichals []Chemichals
}

type Animal string
type Chemichals string

const (
	Cat   Animal = "Cat"
	Manul Animal = "Manul"
	Tiger Animal = "Tiger"

	N  Chemichals = "N"
	Mg Chemichals = "Mg"
	Na Chemichals = "Na"
	K  Chemichals = "K"
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

func WithChemichals(c []Chemichals) func(*Aquarium) {
	return func(a *Aquarium) {
		a.Chemichals = c
	}
}
