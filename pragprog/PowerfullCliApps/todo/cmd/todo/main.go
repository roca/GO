package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/todo"
)

const todoFileName = ".todo.json"

func main() {
	task := flag.String("task", "", "Task to be included in the TODO list")
	list := flag.Bool("list", false, "List all the tasks")
	complete := flag.Int("complete", 0, "Item to be marked as completed")

	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		fmt.Fprintf(os.Stderr, "Invalid option: %s", strings.Join(os.Args[1:], " "))
		os.Exit(1)
	}
}
