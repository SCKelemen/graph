package graph

import (
	"fmt"
)

// Vertex is a node in the graph
type Vertex struct {
	ID      string
	Address string
	Edges   map[string]Edge
}

// Edge is a connection between two vertices
type Edge struct {
	To   Vertex
	From Vertex
}

func NewVertex(name string, address string) Vertex {
	return Vertex{ID: name, Address: address}
}

// AddEdge connect Vertext u to Vertex v
func (u Vertex) AddEdge(v Vertex) error {
	if u.ID == v.ID {
		return VertexLoopException
	}
	for _, edge := range u.Edges {
		if edge.To.ID == v.ID {
			return ParallelEdgeException
		}
	}
	u.Edges[v.ID] = Edge{To: v, From: u}
	fmt.Printf("New Edge: %s -> %s", u.Address, v.Address)
	return nil
}

func (u Vertex) Print() {
	fmt.Println("Vertex")
	fmt.Printf("\tID:\t\t%s\n\tAddress:\t%s\n\tEdges:", u.ID, u.Address)

	for _, edge := range u.Edges {
		fmt.Printf("\t\t%s\t->\t%s", edge.From.ID, edge.To.ID)
		//fmt.Printf("\t\tTo:\t%s\nAddr:\t:%s", edge.To.)
	}
}

// Exception is an error type for this package
type Exception string

// Error returns the error string
func (e Exception) Error() string { return string(e) }

const (
	// VertexLoopException occurs when a Vertex tries to create an edge to itself.
	VertexLoopException Exception = "Vertex cannot loop. Edges cannot connect a vertex to itself. Loops are not allowed. This is not a hypergraph."
	// ParallelEdgeException occurs when a Vertex tries to create multiple edges to the same node.
	ParallelEdgeException Exception = "Vertex cannot have multiple edges with the same vertex. Parallel edges are not allowed. This is not a multigraph."
)
