package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func readline() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // TODO: handle errors
	return sc.Text()
}

func parse(line string) []string {
	return strings.Split(line, " ")
}

func execute(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func loop() {
	for {
		fmt.Print("(^ _ ^) ぎの〜> ")
		line := readline()
		args := parse(line)
		execute(args)
	}
}

func main() {
	loop()
	os.Exit(1)
}
