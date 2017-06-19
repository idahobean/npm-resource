package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/idahobean/npm-resource/check"
	"github.com/idahobean/npm-resource/npm"
)

func main() {

	NPM := npm.NewNPM()
	command := check.NewCommand(NPM)

	var request check.Request
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		fatal("reading request from stdin", err)
	}

	var err error
	if request.Source.PackageName == "" {
		err = errors.New("package_name")
	}
	if err != nil {
		fatal("parameter required", err)
	}

	response, err := command.Run(request)
	if err != nil {
		fatal("running command", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		fatal("writing response to stdout", err)
	}
}

func fatal(message string, err error) {
	fmt.Fprintf(os.Stderr, "error %s: %s\n", message, err)
	os.Exit(1)
}
