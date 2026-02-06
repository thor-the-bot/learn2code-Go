package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: learn2code <command>")
        fmt.Println("commands: run test")
        os.Exit(1)
    }

    cmd := os.Args[1]
    switch cmd {
    case "run":
        // run the exercise's main.go (default hello-world)
        execCmd := exec.Command("go", "run", "./exercises/hello-world/template.go")
        execCmd.Stdout = os.Stdout
        execCmd.Stderr = os.Stderr
        execCmd.Run()
    case "test":
        testCmd := exec.Command("go", "test", "./...", "-v")
        testCmd.Stdout = os.Stdout
        testCmd.Stderr = os.Stderr
        testCmd.Run()
    default:
        fmt.Println("unknown command")
        os.Exit(2)
    }
}
