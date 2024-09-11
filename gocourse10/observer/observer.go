package observer

type Observer interface {
	TimeToEat()
	GetID() int
}

type Item struct {
	ObserverList []Observer
}

func (i *Item) register(o Observer) {
	i.ObserverList = append(i.ObserverList, o)
}

func (i *Item) deregister(o Observer) {
	i.ObserverList = removeFromslice(i.ObserverList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.ObserverList {
		observer.TimeToEat()
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
