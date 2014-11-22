package main

import (
	//"fmt"
	"log"
	"os"
	//"os/exec"
	"github.com/flynn/flynn/Godeps/_workspace/src/github.com/flynn/go-docopt"
)

func init() {
	register("create", runCreate, `
usage: flynn create <name>

Create an application in Flynn.

Examples:

	$ flynn create dsf
	Created dsf
`)
	register("update", runUpdate, `
usage: flynn update <name>

Update an application in Flynn.

Examples:

	$ flynn update dsf
	Updating dsf
`)
	register("deploy", runDeploy, `
usage: flynn deploy [-s <url>]

Options:
	-s, --svn-url <url>  set the svn url of your code

Deploy an application in Flynn.

Examples:

	$ flynn deploy -s http://svnURL
	Exporing svn code
	Compiling code
	Deployed
`)
}

func runCreate(args *docopt.Args) error {
	var appName = args.String["<name>"]

	//exec.Command("git", "remote", "remove", "flynn").Run()
	//exec.Command("git", "remote", "add", "flynn", gitURLPre(clusterConf.GitHost)+app.Name+gitURLSuf).Run()
	os.Setenv("JFLYNN_APP", appName)
	log.Printf("Created %s", appName)
	return nil
}

func runUpdate(args *docopt.Args) error {
	var appName = args.String["<name>"]

	os.Setenv("JFLYNN_APP", appName)
	log.Printf("Start updating %s", os.Getenv("JFLYNN_APP"))
	return nil
}

func runDeploy(args *docopt.Args) error {
	var svn = args.String["--svn-url"]
	log.Printf("Deployed %s", svn)
	return nil
}
