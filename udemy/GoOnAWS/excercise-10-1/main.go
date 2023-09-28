package main

import (
	"excercise-10-1/stacks"
	"flag"
	"fmt"
	"log"
)

func main() {
	stackName := flag.String("s", "", "List resources from a stack by name")
	allStacks := flag.Bool("a", false, "List all stacks")
	flag.Parse()

	if *allStacks {
		stacks, err := stacks.ListStacks(stacks.Client)
		if err != nil {
			log.Fatal(err)
		}
		Print(stacks)
		return
	}

	resources, err := stacks.StackResources(stacks.Client, stackName)
	if err != nil {
		log.Fatalf("Stack with id %s does not exist", *stackName)
	}
	fmt.Println("Resources for stack:", *stackName)
	Print(resources)
}

type X interface {
	stacks.Resource | stacks.Stack
}

func Print[S []E, E X](s S) {
	for _, v := range s {
		fmt.Println(v)
	}
}
