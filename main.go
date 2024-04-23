package main

import (
	// "fmt"
	"log"

	"github.com/BlizzardPurple/DistSys/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	// log.Fatal(tr.ListenAndAccept())
	select {}
	// fmt.Println("Choochie")
}
