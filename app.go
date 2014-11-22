package main

import (
	//"fmt"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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
	register("deploy", runDeploy, `
usage: flynn deploy -a <appName> [-s <url>]

Options:
	-a set app name
	-s, --svn-url <url>  set the svn url of your code

Deploy an application in Flynn.

Examples:

	$ flynn deploy -a dsf -s http://svnURL
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

func runDeploy(args *docopt.Args) error {
	var appName = args.String["-a"]
	var svn = args.String["--svn-url"]
	log.Printf("Exporting %s...", svn)

	var cmd = "docker run -it centos echo haha"
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]
	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)

	cmd = "docker run -it -v /tmp:/tmp -a stdout tegdsf/centos svn export " + svn + " /tmp/" + appName
	log.Println(cmd)
	cmd = "tar cvf /tmp/" + appName + ".tar --directory=/tmp/" + appName + " ."
	cmd = "cat slug.tar | docker run -i -v /tmp/buildpacks:/tmp/buildpacks -e HTTP_SERVER_URL=http://192.168.59.103:8080 -a stdin flynn/slugbuilder - > /tmp/slug.tgz"
	log.Println("Compiling code...")

	log.Printf("Created release for app %s", appName)
	return nil
}
