package main

import{
	"fmt"
	"github.com/BlizzardPurple/DistSys/p2p"
} 

func main() {
	tr:=p2p.NewTCPTransport(":3000")
	log.Fatal(tr.ListenAndAccept())
	select {}
	fmt.Println("Choochie")
}
