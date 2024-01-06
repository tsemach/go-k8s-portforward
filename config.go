package main

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type PFItemPorts struct {
	Src uint16 `yaml:"src"`
	Dst uint16 `yaml:"dst"`
}

type PFItem struct {
	Name      string      `yaml:"name"`
	Pod       string      `yaml:"pod"`
	Namespace string      `yaml:"namespace"`
	Ports     PFItemPorts `yaml:"ports"`
}

type Config []PFItem

var cfg Config

func loadConfig(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(f, &cfg); err != nil {
		log.Fatal(err)
	}
}

func getPFItem(name string) (PFItem, error) {
	for i := 0; i < len(cfg); i++ {
		if name == cfg[i].Name {
			return cfg[i], nil
		}
	}

	return cfg[0], errors.New("unable to find item name")
}
