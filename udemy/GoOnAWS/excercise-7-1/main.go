package main

import (
	"fmt"

	"log"
)

type InstanceState string

const (
	InstanceStateRunning InstanceState = "running"
	InstanceStateStopped InstanceState = "stopped"
)

type Instance struct {
	Name  string
	State InstanceState
}

func LaunchInstance(name string) (*Instance, error) {
	return &Instance{
		Name:  name,
		State: InstanceStateRunning,
	}, nil
}

func StopInstance(instance *Instance) error {
	instance.State = InstanceStateStopped
	return nil
}

func Observe(instance *Instance) {
	fmt.Printf("Instance %s is %v\n", instance.Name, instance.State)	
}

func main() {
	instance, err := LaunchInstance("Alice")
	if err != nil {
		log.Fatalf("LaunchInstance error: %s", err.Error())
	}
	Observe(instance)

	err = StopInstance(instance)
	if err != nil {
		log.Fatalf("StopInstance error: %s", err.Error())
	}
	Observe(instance)
}
