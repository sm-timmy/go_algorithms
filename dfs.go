package main

type Node struct {
	Name     string
	Children []*Node
}

type Graph struct {
	nodes []*Node
}

func (n *Node) dfs(name string) *Node {
	stack := []*Node{n}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Name == name {
			return cur
		}
		for _, child := range cur.Children {
			stack = append(stack, child)
		}
	}
	return nil
}
