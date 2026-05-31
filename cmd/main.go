package main

import (
	"fmt"

	"github.com/birdyhash/gatering"
	"github.com/birdyhash/gatering/gates"
)

/*
this example demonstrates the complete
payload lifecycle.

a payload enters the chamber.

the payload crosses every configured gate.

the transformed payload receives a seal.

the resulting route and hash are displayed.
*/
func main() {

	payload := []byte("hello world")

	chamber := gatering.New()

	chamber.Add(gates.G1)
	chamber.Add(gates.G2)

	result := chamber.Pass(payload)

	seal := gatering.CreateSeal(
		result,
		chamber.Route(),
	)

	fmt.Println("payload :", string(payload))
	fmt.Println("output  :", string(result))
	fmt.Println("")

	fmt.Println("route")

	for _, gate := range seal.Route {
		fmt.Println(" -", gate)
	}

	fmt.Println("")
	fmt.Println("seal")
	fmt.Println(" - hash :", seal.Hash)
	fmt.Println(" - size :", seal.Size)
}
