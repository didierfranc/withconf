package main

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	output, err := exec.Command("go", "run", "main.go", "config.json", "echo", "x").Output()
	if err != nil {
		t.Error(err)
	}

	t.Log(string(output))

	equal := assert.Equal(t, 110, len(output))
	if !equal {
		t.Error("unexpected output")
	}
}
