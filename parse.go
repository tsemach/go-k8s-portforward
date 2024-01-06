package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tsemach/go-k8s-portforward/common"
)

type Color string

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

type Args struct {
	name string
	file string
	help bool
}

func help() {
	colorize(ColorYellow, `
	This tool enable to run multiple kubernetes port forward command. it read configuration file and ack apong it

	config file search:

	1. using -f in the command line
	2. using $PORT_FORWARD_CONFIGFILE environment variable
	3. using pf.yaml | pf.yml in the local directory
	4. using pf.yaml | pf.yml in ~/.config/port-forward

	pf => port forward all found in default config file
	pf -f file.yaml => port forward all found in specfic config file
	pf -n <service-name> => port forward specific name found in default config file
	pf -f file.yaml -n <service-name> => port forward specific name found in specfic config file
`)

	// 	fmt.Printf(`
	// 	This tool enable to run multiple kubernetes port forward command. it read configuration file and ack apong it

	// 	config file search:

	// 	1. using -f in the command line
	// 	2. using $PORT_FORWARD_CONFIGFILE environment variable
	// 	3. using pf.yaml | pf.yml in the local directory
	// 	4. using pf.yaml | pf.yml in ~/.config/port-forward

	//	pf => port forward all found in default config file
	//	pf -f file.yaml => port forward all found in specfic config file
	//	pf -n <service-name> => port forward specific name found in default config file
	//	pf -f file.yaml -n <service-name> => port forward specific name found in specfic config file
	//
	// `)
}

func (args *Args) getFile() string {
	if len(args.file) > 0 {
		return args.file
	}

	if args.file = os.Getenv("PORT_FORWARD_CONFIGFILE"); len(args.file) > 0 {
		return args.file
	}

	if _, err := os.Stat("pf.yaml"); err == nil {
		return "pf.yaml"
	}

	if _, err := os.Stat("pf.yml"); err == nil {
		return "pf.yml"
	}

	return common.First(os.UserHomeDir()) + ".config/port-forward/pf.yaml"
}

func (args *Args) isName() bool {
	return len(args.name) > 0
}

func parse() *Args {
	var args Args

	flag.StringVar(&args.name, "n", "", "a specfic port forward setting to use")
	flag.StringVar(&args.file, "f", "", "a specfic config file to use")
	flag.BoolVar(&args.help, "h", false, "print help message")

	flag.Parse()

	return &args
}
