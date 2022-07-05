package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: demon command [args ...]")
		fmt.Println("Continuously run a program, restart it after it exit. Like a `daemon`")
		return
	}

	prog := os.Args[1]
	args := os.Args[2:]

	log.Printf("[demon]: getting command: %s, with arguments: %v", prog, args)
	for {
		log.Printf("[demon]: [*]--------start--------[*]")

		cmd := exec.Command(prog, args...)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatal(err)
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}

		stdin.Close()
		stdout.Close()
		stderr.Close()

		log.Printf("[demon]: command returned, continue...")
		log.Printf("[demon]: [*]---------end---------[*]")

	}
}
