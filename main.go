package main


func main(){
	root:=Node{"A",nil,nil}
	nodeB:=Node{"B",nil,nil}
	nodeC:=Node{"C",nil,nil}
	nodeD:=Node{"D",nil,nil}
	root.Left=&nodeB
	root.Right=&nodeC
	nodeC.Right=&nodeD
	myTree:=BinaryTree{&root,4}
	ShowTreeGraph(myTree)
}

