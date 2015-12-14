package main

import (
	"flag"
	"fmt"
	"github.com/gracecat/cmdreport/execute"
	"github.com/gracecat/cmdreport/report"
	"io/ioutil"
	"os"
	"strings"
)

var (
	version   = "0.1.0"
	author    = os.Getenv("CMDREPORT_AUTHOR")
	serverURL = os.Getenv("CMDREPORT_SERVER")
)

func main() {
	// create and parse flags
	versionFlag := flag.Bool("v", false, "print the version and exit")
	envFlag := flag.Bool("e", false, "print the env vars and exit")
	flag.Usage = func() {
		fmt.Printf("\ncmdreport is a tool to report stdout/stderr messages to a server.\n\n")
		fmt.Printf("Usage: cmdreport [cmd] | <flag>\n\n")
		fmt.Printf("cmd:\n")
		fmt.Printf("  run a terminal command\n\n")
		fmt.Printf("flag:\n")
		flag.PrintDefaults()
		fmt.Println("")
	}
	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}
	if *envFlag {
		fmt.Println("CMDREPORT_AUTHOR = " + author)
		fmt.Println("CMDREPORT_SERVER = " + serverURL)
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		fmt.Println("missing command...")
		os.Exit(1)
	}
	cmd := os.Args[1]
	cmdArgs := os.Args[2:]
	cmdArgsStr := strings.Join(cmdArgs, " ")
	fmt.Printf("try to run '%s %s'\n", cmd, cmdArgsStr)

	stdout, stderr, err := execute.Execute(cmd, cmdArgs...)
	if err != nil {
		fmt.Println("execution error", err)
		os.Exit(127)
	}

	if len(stdout) != 0 {
		fmt.Println("\nstdout:")
		printStringList(stdout)
	}
	if len(stderr) != 0 {
		fmt.Println("\nstderr:")
		printStringList(stderr)
	}

	// post the report...
	resp, err := report.Post(serverURL, author, cmd, cmdArgs, stdout, stderr)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("server sesponse with status %s\n", resp.Status)
	if resp.StatusCode == 200 {
		fmt.Printf("successful send report (%s)\n", resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}

// pretty print a string array
func printStringList(s []string) {
	for i, d := range s {
		fmt.Printf("%d\t%s\n", i, d)
	}
	fmt.Println("")
}
