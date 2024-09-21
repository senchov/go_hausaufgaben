package aquarium

import "testing"

func TestNew(t *testing.T) {
	a := New()

	if a.Animal != "" && a.Size != 0 && a.SaltLevel != 0 {
		t.Errorf("Aquarium created with not default values, was animal=%s, size=%d, salt=%d", a.Animal, a.Size, a.SaltLevel)
	}

	a = New(WithAnimal(Cat))

	if a.Animal != Cat {
		t.Errorf("Aquarium created with wrong animal should be cat but was %s", a.Animal)
	}

	a = New(
		WithAnimal(Manul),
		WithSaltLevel(5),
		WithSize(10),
	)

	if a.Animal != Manul || a.SaltLevel != 5 || a.Size != 10 {
		t.Errorf("Aquarium created with wrong values should be Animal Manul , saltLevel 6,size 10 but was %s saltLevel %d size %d", a.Animal, a.SaltLevel, a.Size)
	}
}
