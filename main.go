package main

import (
	"context"
	"log"
	"time"

	// "github.com/justinbarrick/go-k8s-portforward"
	"github.com/justinbarrick/go-k8s-portforward/portforward"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/**
pf => port forward all found in default config file
pf -f file.yaml => port forward all found in specfic config file
pf -n <service-name> => port forward specific name found in default config file
pf -f file.yaml -n <service-name> => port forward specific name found in specfic config file
*/

func main() {
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
	time.Sleep(600 * time.Second)
}
