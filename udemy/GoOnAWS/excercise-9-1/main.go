package main

import (
	"excercise-9-1/instanceLister"
	"fmt"
)

func main() {
	instances, err := instanceLister.ListInstances(instanceLister.Client)
	if err != nil {
		fmt.Println("Error with listing instances: ", err)
	}

	for _, instance := range instances {
		fmt.Printf("Instance %s is %s\n", instance.Name, instance.State)
	}
}
