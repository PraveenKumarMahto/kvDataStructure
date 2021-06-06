package bst

import "kvDataStructure/utility"

type Node struct {
	Key   *int64
	Value *string
	Left  *Node
	Right *Node
}

func createNode(key int64, value string) *Node {
	return &Node{
		Key:   &key,
		Value: &value,
	}
}

func (root *Node) InsertOrUpdateNode(key int64, value string) *Node {

	if root == nil {
		return createNode(key, value)
	}
	start := root
	for {
		if root.Key == &key {
			root.Value = &value
			break
		}
		if *root.Key < key {
			if root.Right != nil {
				root = root.Right
				continue
			}
			root.Right = createNode(key, value)
			break
		}
		if root.Left != nil {
			root = root.Left
			continue
		}
		root.Left = createNode(key, value)
		break
	}
	return start
}

func (root *Node) SearchNode(key int64) *string {
	for root != nil {
		if root.Key == &key {
			return root.Value
		}
		if *root.Key < key {
			root = root.Right
			continue
		}
		root = root.Left
	}
	return nil
}

func (root *Node) Display() {
	if root != nil {
		utility.PrintKeyValue(*root.Key, *root.Value)
		root.Left.Display()
		root.Right.Display()
	}
}
