package main

type BinaryTree struct {
	Root *Node
	NumNodes int
}

type Node struct {
	Value string
	Left *Node
	Right *Node
}

type nodePair struct {
	ValA, ValB string
}

type nodePos struct {
	Val string
	YPos int
	XPos int
}

var data []nodePos
var endPoints []nodePair


