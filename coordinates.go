package main

import "fmt"

func prepareDrawTree(tree BinaryTree){
	prepareToDraw(tree)
	fmt.Println(endPoints)
	fmt.Println(data)
}

func findXY(val string)(int, int){
	for i:=0;i<len(data);i++{
		if data[i].Val == val{
			return data[i].XPos,data[i].YPos
		}
	}
	return -1,-1
}

func findX(val string)int{
	for i:=0;i<len(data);i++{
		if data[i].Val == val{
			return i
		}
	}
	return -1
}

func setXValues(){
	for index:=0;index<len(data);index++{
		xValue:=findX(data[index].Val)
		data[index].XPos=xValue
	}
}

func prepareToDraw(tree BinaryTree){
	inorderLevel(tree.Root,1)
	setXValues()
	getEndPoints(tree.Root,nil)
}

func inorderLevel(node *Node,level int){
	if node != nil {
		inorderLevel(node.Left,level+1)
		data = append(data,nodePos{node.Value,100-level,-1})
		inorderLevel(node.Right,level+1)
	}
}

func getEndPoints(node *Node, parent *Node){
	if node != nil {
		if parent != nil {
			endPoints = append(endPoints,nodePair{node.Value,parent.Value})
		}
		getEndPoints(node.Left,node)
		getEndPoints(node.Right,node)
	}
}


