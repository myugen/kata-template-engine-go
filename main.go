package main

import (
	"fmt"

	"github.com/myugen/kata-template-engine-go/kata"
)

func main() {
	dependency := kata.NewDependencyImpl()
	k := kata.New(dependency)

	fmt.Printf("the value is %d", k.Value())
}
