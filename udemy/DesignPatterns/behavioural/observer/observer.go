package observer

type IObserver interface {
	Notify(string)
}

type IPublisher interface {
	AddObserver(o IObserver)
	RemoveObserver(o IObserver)
	NotifyObservers(m string)
}

type Publisher struct {
	ObserversList []IObserver
}

func (s *Publisher) AddObserver(o IObserver) {
	s.ObserversList = append(s.ObserversList, o)
}
func (s *Publisher) RemoveObserver(o IObserver) {
	var indexToRemove int
	for i, observer := range s.ObserversList {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	s.ObserversList = append(s.ObserversList[:indexToRemove], s.ObserversList[indexToRemove+1:]...)
}
func (s *Publisher) NotifyObservers(m string) {
	for _, observer := range s.ObserversList {
		observer.Notify(m)
	}
}
