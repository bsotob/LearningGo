//binaryTree its a simple exemple how can be implemented a binary tree
//in golang, but its necessary that you understand that this is not a
//BST tree, because here its not contemplated the position of a node
//depending on its value.
package main

import "fmt"

//Node in a binary Tree
type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

//BinaryTree is a data structure to store data with a tree idea
//cant be repeated value of nodes.
type BinaryTree struct {
	Root      *Node
	RootValue int64
	Height    int64
}

//maxNumber returns the max number of the pair x, y
func maxNumber(x, y int64) int64 {
	if x >= y {
		return x
	}
	return y
}

//GetTreeHeight returns the height of the tree
func GetTreeHeight(p *Node) int64 {
	if p != nil {
		var (
			left, right int64
		)
		if p.Left != nil {
			left = GetTreeHeight(p.Left) + 1
		}
		if p.Right != nil {
			right = GetTreeHeight(p.Right) + 1
		}
		return maxNumber(left, right)
	}
	return -1
}

//IsPresent returns true if x is the tree pointed by p
func IsPresent(x int64, p *Node) bool {
	if p != nil {
		if p.Value == x {
			return true
		}
		if p.Left != nil {
			return IsPresent(x, p.Left)
		}
		if p.Right != nil {
			return IsPresent(x, p.Right)
		}
	}
	return false
}

//DeleteSubTree delete receives the value of the node that wants to
//be deleted, this will delete all the substrees if the node was found
//and its not the actual root of the BinaryTree
func DeleteSubTree(p *Node, rootValue int64, value int64) {
	//TODO: MAKE THIS WORK
	if p != nil {
		if p.Value != rootValue && p.Value == value {
			p = &Node{}
		} else if p.Left != nil {
			DeleteSubTree(p.Left, rootValue, value)
		} else if p.Right != nil {
			DeleteSubTree(p.Right, rootValue, value)
		}
	}
}

//PrintTree prints the BinaryTree with the next logics:
//the first printed node its the root of the BinaryTree, then starts
//walking in the left subtree, a leaf it's identified when its follow
//by two 0's, soo a 3 0 0, 3 is a leaf, 10 3 0 4 0 0 its:
//     10
//    /
//   3
//  / \
// 0    4
//     / \
//    0   0
func PrintTree(p *Node) {
	if p != nil {
		fmt.Printf("%v\n", p.Value)
		if p.Left != nil {
			PrintTree(p.Left)
		}
		if p.Left == nil {
			fmt.Printf("0\n")
		}
		if p.Right != nil {
			PrintTree(p.Right)
		}
		if p.Right == nil {
			fmt.Printf("0\n")
		}
	}
}

func main() {
	//In this exemple you cant insert a node because, theres any rule
	//that tells us where to insert the node.
	n1right := &Node{4, nil, nil}
	n1 := &Node{3, nil, n1right}
	n2 := &Node{11, nil, nil}
	root := &Node{10, n1, n2}
	b := &BinaryTree{root, root.Value, GetTreeHeight(root)}

	fmt.Println(IsPresent(2, b.Root))
	fmt.Println(b)
	PrintTree(b.Root)
	//DeleteSubTree(b.Root, b.RootValue, 4)
	//fmt.Println("-----------------------")
	//PrintTree(b.Root)
}
