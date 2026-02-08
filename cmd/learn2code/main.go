package main

import (
	"flag"
	"fmt"
)

func main() {
	list := flag.Bool("list", false, "List available exercises")
	run := flag.String("run", "", "Run an exercise by name")
	flag.Parse()

	if *list {
		fmt.Println("Available exercises:\n- hello-world")
		return
	}
	if *run != "" {
		fmt.Printf("Running exercise: %s\n", *run)
		// TODO: implement exercise runner
		return
	}
	fmt.Println("learn2code CLI — use -list or -run <exercise>")
}
