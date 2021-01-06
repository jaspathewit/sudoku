package grid

import "fmt"

// methods implemented by a topology
type Topology interface {
	// create a new grid in this topology
	NewGrid() (*Grid, error)
	// string representation of a grid in this topology
	AsString(g *Grid) string
}

// a type for maintaining a slice of cell references.
type Peers []string

// a type for a slice of peers
type PeerSet []Peers

// function locates the peers for a cell reference
func (ps PeerSet) FindPeersFor(ref string) (Peers, error) {
	for _, peers := range ps {
		for _, r := range peers {
			if r == ref {
				return peers, nil
			}
		}
	}

	return nil, fmt.Errorf("something wrong could not find peers for %s", ref)
}