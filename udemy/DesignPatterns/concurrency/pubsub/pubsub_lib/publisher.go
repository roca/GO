package pubsub_lib

type IPublisher interface {
	start()
	AddSubscriberCh() chan<- ISubscriber
	RemoveSubscriberCh() chan<- ISubscriber
	PublishingCh() chan<- interface{}
	Stop()
}

type publisher struct {
	subscribers []ISubscriber
	addSubCh    chan ISubscriber
	removeSubCh chan ISubscriber
	in          chan interface{}
	stop        chan struct{}
}

func (p *publisher) start() {
	for {
		select {
		case msg := <-p.in:
			for _, sub := range p.subscribers {
				sub.Notify(msg)
			}
		case sub := <-p.addSubCh:
			p.subscribers = append(p.subscribers, sub)
		case sub := <-p.removeSubCh:
			for i, candidate := range p.subscribers {
				if candidate == sub {
					p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
					candidate.Close()
					break
				}
			}
		case <-p.stop:
			for _, sub := range p.subscribers {
				sub.Close()
			}
			close(p.addSubCh)
			close(p.in)
			close(p.removeSubCh)
			return

		}
	}
}

func (p *publisher) AddSubscriberCh() chan<- ISubscriber {
	return p.addSubCh
}
func (p *publisher) RemoveSubscriberCh() chan<- ISubscriber {
	return p.removeSubCh
}
func (p *publisher) PublishingCh() chan<- interface{} {
	return p.in
}
func (p *publisher) Stop() {
	close(p.stop)
}

func NewPublisher() IPublisher {
	p := &publisher{
		//subscribers: []ISubscriber{},
		addSubCh:    make(chan ISubscriber),
		removeSubCh: make(chan ISubscriber),
		in:          make(chan interface{}),
		stop:        make(chan struct{}),
	}

	go p.start()

	return p
}
