package gtable

import "fmt"

type Node struct {
	Peers map[string]Peer
}

func (n Node) Broadcast(message string) int {
	count := 0
	for _, p := range n.Peers {
		fmt.Println(p.Address)
		count++
	}
	return count
}

type Peer struct {
	Address string
	ID      string
}
