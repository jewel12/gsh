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

// simple parser
func parse(line string) []*exec.Cmd {
	var cmds []*exec.Cmd
	cmdexps := strings.Split(line, "|")
	for _, ce := range cmdexps {
		var args []string
		argsexps := strings.Split(ce, " ")
		for _, a := range argsexps {
			// trim white spaces: strings.Split(" gino wan ") => ["", "gino", "wan", ""]
			if a == "" {
				continue
			}
			args = append(args, a)
		}
		cmds = append(cmds, exec.Command(args[0], args[1:]...))
	}
	return cmds
}

func execute(cmds []*exec.Cmd) error {
	pipe(cmds)
	for _, c := range cmds {
		c.Start()
	}
	for _, c := range cmds {
		if err := c.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func pipe(cmds []*exec.Cmd) {
	n := len(cmds)

	if n == 1 {
		c := cmds[0]
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return
	}

	for i, c := range cmds {
		c.Stderr = os.Stderr // redirect all process stderrs to parent's stderr
		switch i {
		case 0: // first command
			c.Stdin = os.Stdin
			cmds[i+1].Stdin, _ = c.StdoutPipe()
		case (n - 1): // last command
			c.Stdout = os.Stdout
		default:
			cmds[i+1].Stdin, _ = c.StdoutPipe()
		}
	}
}

func loop() {
	for {
		fmt.Print("GINO (^ _ ^) WAN >> ")
		line := readline()
		cmds := parse(line)
		execute(cmds)
	}
}

func main() {
	loop()
	os.Exit(1)
}
