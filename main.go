package main

import (
	"github.com/tanvircs/afsa/cmd"
)

func init() {
	// Ensure colored output on Windows terminals
	if cmd.ForceColor {
		_ = cmd.SetupColor()
	}
}

func main() {
	cmd.Execute()
}
