package main

import (
	"fmt"
	"reflect"
)

type node struct {
	data        int
	right, left *node
}

type pair struct {
	height   int
	diameter int
}

type balancedPair struct {
	height   int
	balanced bool
}

func buildTree(root *node) *node {
	//fmt.Println("Enter node value")
	var i int
	var _, _ = fmt.Scanf("%d", &i)
	if i == -1 {
		return root
	}

	if root == nil {
		root = new(node)
	}
	root.data = i
	root.left = buildTree(root.left)
	root.right = buildTree(root.right)
	return root
}

func height(root *node) int {
	if root == nil {
		return 0
	}
	var heightRight = height(root.right)
	var heightLeft = height(root.left)

	return Max(heightRight, heightLeft) + 1
}

func (rootWithTree node) IsEmpty() bool {
	return reflect.DeepEqual(rootWithTree, node{})
}

func printTreePreOrder(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	printTreePreOrder(root.left)
	printTreePreOrder(root.right)
}

func printTreeInorder(root *node) {
	if root == nil {
		return
	}
	printTreePreOrder(root.left)
	fmt.Println(root.data)
	printTreePreOrder(root.right)
}

func printTreePostorder(root *node) {
	if root == nil {
		return
	}
	printTreePostorder(root.left)
	printTreePostorder(root.right)
	fmt.Println(root.data)
}

func printKthLevel(root *node, level int) {
	if root == nil {
		return
	}
	if level == 1 {
		fmt.Print(" ", root.data)
	}
	printKthLevel(root.left, level-1)
	printKthLevel(root.right, level-1)
}

func levelOrderTraversal(root *node) {
	var height = height(root)
	for i := 1; i <= height; i++ {
		printKthLevel(root, i)
		fmt.Println()
	}
}

func BFS(root *node) {
	var q []*node
	q = append(q, root)
	for len(q) != 0 {
		el := q[0]
		q = q[1:]
		fmt.Print(" ", el.data)
		if el.left != nil {
			q = append(q, el.left)
		}
		if el.right != nil {
			q = append(q, el.right)
		}
	}
}

func BFS_LevelByLevel(root *node) {
	var q []*node
	q = append(q, root)
	q = append(q, nil)
	for len(q) != 0 {
		el := q[0]
		if el == nil {
			q = q[1:]
			fmt.Println()
			if len(q) != 0 {
				q = append(q, nil)
			}
		} else {
			q = q[1:]
			fmt.Print(" ", el.data)
			if el.left != nil {
				q = append(q, el.left)
			}
			if el.right != nil {
				q = append(q, el.right)
			}
		}
	}
}

func count(root *node) int {
	if root == nil {
		return 0
	}
	return 1 + count(root.left) + count(root.right)
}

func sum(root *node) int {
	if root == nil {
		return 0
	}
	return sum(root.left) + sum(root.right) + root.data
}

func diameter(root *node) int {
	if root == nil {
		return 0
	}
	h1 := height(root.left)
	h2 := height(root.right)
	d1 := h1 + h2
	d2 := diameter(root.left)
	d3 := diameter(root.right)
	return Max(d1, Max(d2, d3))
}

func fastDiameter(root *node) *pair {
	p := new(pair)
	if root == nil {
		p.diameter = 0
		p.height = 0
		return p
	}
	rightPair := fastDiameter(root.right)
	leftPair := fastDiameter(root.left)

	p.height = Max(leftPair.height, rightPair.height) + 1
	p.diameter = Max(rightPair.height+leftPair.height, Max(rightPair.diameter, leftPair.diameter))
	return p
}

func isHeightBalanced(root *node) *balancedPair {
	p := new(balancedPair)
	if root == nil {
		p.balanced = true
		p.height = 0
		return p
	}

	left := isHeightBalanced(root.left)
	right := isHeightBalanced(root.right)

	p.height = Max(left.height, right.height) + 1

	if left.balanced && right.balanced && Abs(left.height-right.height) <= 2 {
		p.balanced = true
	}else {
		p.balanced = false
	}
	return p;
}



func main() {
	fmt.Println("Building a binary tree")
	var root *node
	root = buildTree(root)
	fmt.Println("Printing the tree pre-order:")
	printTreePreOrder(root)
	fmt.Println("Printing the tree in-order:")
	printTreeInorder(root)
	fmt.Println("Printing the tree post-order:")
	printTreePostorder(root)
	//fmt.Println("Height of the tree: %d", height(root))
	var heightOfTree = height(root)
	fmt.Println("Height of tree is ", heightOfTree)
	fmt.Println("Going to print tree level-wise")
	levelOrderTraversal(root)
	fmt.Println("Going to print tree bfs level-wise")
	BFS_LevelByLevel(root)
	fmt.Println("Total count of nodes is", count(root))
	fmt.Println("Sum of all the node is ", sum(root))
	fmt.Println("Diameter of tree is ", diameter(root))
	fmt.Println("Fast diameter of tree is ", fastDiameter(root).diameter)
	fmt.Println("Is the tree balanced ", isHeightBalanced(root).balanced)
}
