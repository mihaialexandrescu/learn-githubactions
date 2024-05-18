package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"sigs.k8s.io/yaml"

	"github.com/mihaialexandrescu/learn-githubactions/gomodule01/nothing"
)

func main() {
	fmt.Println(nothing.Hello)
	spew.Dump(nothing.Hello)
	_ = yaml.YAMLToJSON
}
