package main

import (
	"github.com/duckbrain/shiboleet/cmd"
	"github.com/duckbrain/shiboleet/lib/runner"
)

//go:generate go run ./cmd/sqlboiler
//go:generate go run ./cmd/gqlgen

func main() {
	runner.Main(cmd.RootCmd)
}
