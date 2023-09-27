package main

import (
	"excercise-10-1/stacks"
	"flag"
	"fmt"
)

func main() {
	stackName := flag.String("stack_name", "", "List resources from a stack")
	allStacks := flag.Bool("all", false, "List all stacks")
	flag.Parse()

	if *allStacks {
		stacks, err := stacks.ListStacks(stacks.Client)
		if err != nil {
			panic(err)
		}
		Print(stacks)
	} else {

		fmt.Println("Resources for stack:", *stackName)
		resources, err := stacks.ListStackResources(stacks.Client, *stackName)
		if err != nil {
			panic(err)
		}
		Print(resources)
	}
}

type X interface {
	stacks.Resource | stacks.Stack
}

func Print[S []E, E X](s S) {
	for _, v := range s {
		fmt.Println(v)
	}
}
