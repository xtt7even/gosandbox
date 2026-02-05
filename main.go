package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
)

type peer struct {
	id string
	connectionRadius int 
	name string
	maxShareSpeedPerSec float32	
}

func newPeer(name string) *peer {
	id := uuid.New();
	randomSpeed := rand.IntN(50)
	randomRadius := rand.IntN(10)
	if randomSpeed < 10 {
		randomSpeed = 10
	}
	maxShareSpeedPerSec := float32(randomSpeed)
	connectionRadius := randomRadius
	p := peer{id: id.String(), name: name, connectionRadius: connectionRadius, maxShareSpeedPerSec: maxShareSpeedPerSec }
	return &p
}

func main() {
	peer := newPeer("First peer")
	fmt.Println("New peer found:")
	fmt.Println("Peer name:", peer.name)
	fmt.Println("Peer id:", peer.id)
	fmt.Println("Peer max share speed (KB/S):", peer.maxShareSpeedPerSec)
	fmt.Println("Peer availability radius (M):", peer.connectionRadius)

}