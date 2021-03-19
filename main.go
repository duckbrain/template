package main

import (
	"github.com/duckbrain/shiboleet/cmd"
	"github.com/duckbrain/shiboleet/lib/runner"
)

func main() {
	runner.Main(cmd.RootCmd)
}
