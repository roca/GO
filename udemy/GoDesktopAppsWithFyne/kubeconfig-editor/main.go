package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig string
)

func main() {
	defaultPath := fmt.Sprintf("%s/.kube/config",os.Getenv("HOME"))
	flag.StringVar(&kubeconfig, "kubeconfig",  defaultPath, "Path to kube config file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}
