package channel_singleton

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

func init() {
	var count int
	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				break
			}
		}
	}(addCh, getCountCh, quitCh)
}

type ISingleton interface {
	AddOne()
	GetCount() int
	Stop()
}

type singleton struct {
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

var instance *singleton

func GetInstance() ISingleton { // Return object must match interface
	if instance == nil {
		instance = new(singleton) // new() returns a pointer
	}
	return instance
}
