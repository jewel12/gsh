package main

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cmdtests := []struct {
		line string
		cmds []*exec.Cmd
	}{
		{"gino", []*exec.Cmd{exec.Command("gino")}},
		{"gino | wan -s", []*exec.Cmd{exec.Command("gino"), exec.Command("wan", "-s")}},
		{"gino   | wan  -s", []*exec.Cmd{exec.Command("gino"), exec.Command("wan", "-s")}},
	}
	for _, tt := range cmdtests {
		t.Run(tt.line, func(t *testing.T) {
			cmds := parse(tt.line)
			assert.Equal(t, tt.cmds, cmds)
		})
	}
}
