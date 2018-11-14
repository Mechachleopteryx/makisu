package main

import (
	"log"
	"os"

	"github.com/uber/makisu/cli"
	"github.com/uber/makisu/lib/utils"

	"github.com/apourchet/commander"
)

func main() {
	log.Printf("Starting makisu (version=%s)", utils.BuildHash)

	application := cli.NewClientApplication()
	cmd := commander.New()
	if err := cmd.RunCLI(application, os.Args[1:]); err != nil {
		log.Fatalf("Command failure: %s", err)
	}
}