package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func addNode(node *Node, r Record) *Node {
	if node.ID == r.Parent {
		node.Children = append(node.Children, &Node{ID: r.ID})
		return node
	}

	if len(node.Children) > 0 {
		for _, child := range node.Children {
			addNode(child, r)
		}
	}

	return nil
}

func Build(records []Record) (*Node, error) {

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if !inputIsVaid(records) {
		return nil, errors.New("Erro")
	}

	if len(records) == 0 {
		return nil, nil
	} else if len(records) == 1 {
		return &Node{ID: 0}, nil
	}

	n := &Node{ID: 0}

	for _, r := range records[1:] {
		addNode(n, r)
	}

	return n, nil
}

func inputIsVaid(records []Record) bool {

	for i, r := range records {

		if r.ID != i || r.Parent > r.ID || (r.ID > 0 && r.Parent == r.ID) {
			return false
		}
	}
	return true
}
