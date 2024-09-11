package observer

import "testing"

func TestNotifyAll(t *testing.T) {
	item := Item{}
	testObserver := TestObserver{}
	item.Register(&testObserver)
	item.NotifyAll()

	if !testObserver.value {
		t.Error("Event wasn't received")
	}

	item.Deregister(&testObserver)
	if !testObserver.value {
		t.Errorf("Event was received but shouldn't")
	}
}

type TestObserver struct {
	value bool
}

func (t *TestObserver) TimeToEat() {
	t.value = !t.value
}

func (c TestObserver) GetID() int {
	return 1
}
