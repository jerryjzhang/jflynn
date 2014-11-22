package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/flynn/flynn/Godeps/_workspace/src/github.com/flynn/go-docopt"
)

var (
	flagCluster = os.Getenv("FLYNN_CLUSTER")
	flagApp     string
)

func main() {
	log.SetFlags(0)

	usage := `
usage: flynn [-a <app>] <command> [<args>...]

Options:
	-a <app>
	-h, --help

Commands:
	help      show usage for a specific command
	cluster   manage clusters
	create    create an app
	delete    delete an app
	apps      list apps
	ps        list jobs
	kill      kill a job
	log       get job log
	scale     change formation
	run       run a job
	env       manage env variables
	route     manage routes
	provider  manage resource providers
	resource  provision a new resource
	key       manage SSH public keys
	release   add a docker image release
	version   show flynn version

See 'flynn help <command>' for more information on a specific command.
`[1:]
	args, _ := docopt.Parse(usage, nil, true, Version, true)

	cmd := args.String["<command>"]
	cmdArgs := args.All["<args>"].([]string)

	if cmd == "help" {
		if len(cmdArgs) == 0 { // `flynn help`
			fmt.Println(usage)
			return
		} else { // `flynn help <command>`
			cmd = cmdArgs[0]
			cmdArgs = make([]string, 1)
			cmdArgs[0] = "--help"
		}
	}

	if err := runCommand(cmd, cmdArgs); err != nil {
		log.Fatal(err)
		return
	}
}

type command struct {
	usage     string
	f         interface{}
	optsFirst bool
}

var commands = make(map[string]*command)

func register(cmd string, f interface{}, usage string) *command {
	switch f.(type) {
	case func(*docopt.Args) error, func() error, func():
	default:
		panic(fmt.Sprintf("invalid command function %s '%T'", cmd, f))
	}
	c := &command{usage: strings.TrimLeftFunc(usage, unicode.IsSpace), f: f}
	commands[cmd] = c
	return c
}

func runCommand(name string, args []string) (err error) {
	argv := make([]string, 1, 1+len(args))
	argv[0] = name
	argv = append(argv, args...)

	cmd, ok := commands[name]
	if !ok {
		return fmt.Errorf("%s is not a flynn command. See 'flynn help'", name)
	}
	parsedArgs, err := docopt.Parse(cmd.usage, argv, true, "", cmd.optsFirst)
	if err != nil {
		return err
	}

	switch f := cmd.f.(type) {
	case func(*docopt.Args) error:
		return f(parsedArgs)
	case func() error:
		return f()
	case func():
		f()
		return nil
	}

	return fmt.Errorf("unexpected command type %T", cmd.f)
}
