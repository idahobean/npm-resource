package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/idahobean/npm-resource/npm"
	"github.com/idahobean/npm-resource/out"
)

func main() {

	NPM := out.NewNPM()
	command := out.NewCommand(NPM)

	var request out.Request
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		fatal("reading request from stdin", err)
	}

	var err error
	if request.Params.Path == "" {
		err = errors.New("path")
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
