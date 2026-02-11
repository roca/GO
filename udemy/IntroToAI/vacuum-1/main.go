package main

import "flag"

func main() {
	var configFile, algorithm string
	var animate bool
	flag.StringVar(&configFile, "file", "empty.json", "Path to the configuration file")
	flag.StringVar(&algorithm, "algorithm", "random", "Cleaning Algorithm, Cleaning algorithm to use (random, ...")
	flag.BoolVar(&animate, "animate", true, "Animate while cleaning")
	flag.Parse()

	room := NewRoom(configFile, animate)
	_ = room

}
