package memento

import "errors"

type State struct {
	Description string
}
type memento struct {
	state State
}
type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}
func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	// Does nothing
	c.mementoList = append(c.mementoList, m)
}
func (c *careTaker) Memento(i int) (memento, error) {
	if len(c.mementoList) != 0 && i >= 0 {
		return c.mementoList[i],nil
	}
	return memento{}, errors.New("State noy found")
}
