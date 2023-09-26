package main

import (
	"excercise-10-1/stacks"
	"flag"
	"fmt"
)

func main() {
	stackName := flag.String("stack_name", "", "Stack name")
	flag.Parse()

	fmt.Println(*stackName)

	resources, err := stacks.ListStackResources(stacks.Client, *stackName)
	if err != nil {
		panic(err)
	}

	for _, resource := range resources {
		fmt.Println(resource.LogicalID, resource.Status)
	}
}
