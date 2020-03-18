package trees

//RbTree red black tree go implementation
//based off my implementation in python
type RbTree struct {
	Root *Node
	size int
}

//Len gets length of RbTree
func Len(tree *RbTree) int {
	return tree.size
}

//InitializeTree initializes tree since it uses dummy nodes as leafs
//if Node.Entry == nil then its a dummy node
func InitializeTree() *RbTree {
	t := new(RbTree)
	t.Root = MakeNode(false)
	return t
}

// IsLeaf returns true if node is a leaf
func IsLeaf(n *Node) bool {
	return n.Entry == nil
}

//GetSibling gets sibling of node
func GetSibling(n *Node) *Node {
	if n.parent != nil {
		if n == n.parent.right {
			return n.parent.left
		}
		return n.parent.right
	}
	return nil
}

//GetMaxOfTree returns max value of tree at root
func GetMaxOfTree(r *Node) *Node {
	if IsLeaf(r.right) {
		return r
	}
	return GetMaxOfTree(r.right)
}

//GetMinOfTree returns min value of tree at root
func GetMinOfTree(r *Node) *Node {
	if IsLeaf(r.left) {
		return r
	}
	return GetMinOfTree(r.left)
}

//function to locate key in tree, returns both the node with the key and the parent node
//(to make restructuring easier)
func seek(key int, curr *Node, par *Node) (*Node, *Node) {
	if IsLeaf(curr) {
		return curr, par
	}
	if key == curr.Entry.Key {
		return curr, par
	} else if key < curr.Entry.Key {
		return seek(key, curr.left, curr)
	} else {
		return seek(key, curr.right, curr)
	}
}

//Get retrieves associated value with key
func (r *RbTree) Get(key int) interface{} {
	n, _ := seek(key, r.Root, nil)
	if IsLeaf(n) {
		return nil
	}
	return n.Entry.Value
}

//satisfies rb tree properties after putting new value in
func remedyPut(n *Node) {

}

//fixes a double black node
//black == false red == true
func fixDoubleBlack(n *Node) {

}

//satisfies rb tree properties after removing a node
func remedyRemove(n *Node) {

}
