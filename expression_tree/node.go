package expression_tree

import "github.com/Qwertymart/Roman_Calculator/roman"

type Node struct {
	Op    string
	Left  *Node
	Right *Node
	Value *roman.Numeral
}

func NewValueNode(val *roman.Numeral) *Node {
	return &Node{Value: val}
}

func NewOpNode(op string, left *Node, right *Node) *Node {
	return &Node{Op: op, Left: left, Right: right}
}
