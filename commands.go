package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandCreate,
	commandUpdate,
	commandDeploy,
	commandScale,
}

var commandCreate = cli.Command{
	Name:  "create",
	Usage: "jflynn create <appName>",
	Description: `
`,
	Action: doCreate,
}

var commandUpdate = cli.Command{
	Name:  "update",
	Usage: "",
	Description: `
`,
	Action: doUpdate,
}

var commandDeploy = cli.Command{
	Name:  "deploy",
	Usage: "",
	Description: `
`,
	Action: doDeploy,
}

var commandScale = cli.Command{
	Name:  "scale",
	Usage: "",
	Description: `
`,
	Action: doScale,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doCreate(c *cli.Context) {
	log.Println("Created app successfully")
}

func doUpdate(c *cli.Context) {
}

func doDeploy(c *cli.Context) {
}

func doScale(c *cli.Context) {
}
