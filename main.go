package main

import (
	"fmt"
	"log"

	"https://github.com/BlizzardPurple/DistSys/tree/master/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	log.Fatal(tr.ListenAndAccept())
	select {}
	fmt.Println("Choochie")
}
