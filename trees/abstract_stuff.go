package trees

//Node n
type Node struct {
	Entry  *Entry
	Color  bool
	parent *Node
	left   *Node
	right  *Node
}

//Entry e
type Entry struct {
	Key   int
	Value interface{}
}

//MakeNode returns pointer to new node
func MakeNode(color bool) *Node {
	n := new(Node)
	n.Color = color
	return n
}
