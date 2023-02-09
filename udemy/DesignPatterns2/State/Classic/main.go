package main

import "fmt"

// State Design Pattern

// Here is a simple state machine with just two states.

type Switch struct {
	State State
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

func NewSwitch() *Switch {
	return &Switch{State: NewOffState()}
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct{}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}
func (b *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState: BaseState{}}
}

func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning light off...")
	sw.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState: BaseState{}}
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on...")
	sw.State = NewOnState()
}

func main() {
	sw := NewSwitch()
	sw.On()
	sw.On()
	sw.Off()
	sw.Off()
}
