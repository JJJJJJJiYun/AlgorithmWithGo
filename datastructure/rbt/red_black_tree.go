package rbt

type NodeColor bool

type Node struct {
	val    int
	color  NodeColor
	left   *Node
	right  *Node
	parent *Node
}

const (
	NodeColorRed   = NodeColor(false)
	NodeColorBlack = NodeColor(true)
)
