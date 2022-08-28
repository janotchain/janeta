package main

import (
	"runtime"

	cmd "github.com/janotchain/janeta/cmd/janetacli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
