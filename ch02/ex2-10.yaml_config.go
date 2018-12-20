// Listing 2.10  Parsing a YAML configuration file: yaml_config.go

package main

import (
	"fmt"
	"github.com/kylelemons/go-gypay/yaml"
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}
