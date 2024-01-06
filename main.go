package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tsemach/go-k8s-portforward/portforward"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/**
config file search:
	1. using -f in the command line
	2. using $PORT_FORWARD_CONFIGFILE environment variable
	3. using pf.yaml | pf.yml in the local directory
	4. using pf.yaml | pf.yml in ~/.config/port-forward

pf => port forward all found in default config file
pf -f file.yaml => port forward all found in specfic config file
pf -n <service-name> => port forward specific name found in default config file
pf -f file.yaml -n <service-name> => port forward specific name found in specfic config file
*/

func mainOld() {
	var context = context.Background()

	pf, err := portforward.NewPortForwarder("httpd", metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app":   "httpd",
			"order": "2",
		},
	}, 80)
	if err != nil {
		log.Fatal("Error setting up port forwarder: ", err)
	}

	err = pf.Start(context)
	if err != nil {
		log.Fatal("Error starting port forward: ", err)
	}

	log.Printf("Started tunnel on %d\n", pf.ListenPort)
	// for true {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Printf("ffor loop")
	// }	

	
	time.Sleep(600 * time.Second)
}

func handleItem(name string) {
	// item, err := getPFItem(name)
	// if err != nil {
	// 	log.Fatal("unable to find name in config file")
	// 	return 
	// }


}

func main() {
	var args = parse()
	loadConfig(args.getFile())

	fmt.Println("name:", args.name)
	fmt.Println("file:", args.getFile())

	if args.isName() {
		handleItem(args.name)

		return
	}
}
