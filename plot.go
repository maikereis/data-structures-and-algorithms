package main

import (
	"image/color"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func drawGraph(a fyne.App, w fyne.Window, path string) {
	image := canvas.NewImageFromResource(theme.FyneLogo())
	image = canvas.NewImageFromFile(path + "tree.png")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	w.Show()
}

func ShowTreeGraph(myTree BinaryTree) {
	prepareDrawTree(myTree)
	myApp := app.New()
	myWindow := myApp.NewWindow("BinaryTree")

	// Get current working directory (project folder)
	currentDir, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	path := currentDir + "/"

	nodePts := make(plotter.XYs, myTree.NumNodes)
	for i := 0; i < len(data); i++ {
		nodePts[i].Y = float64(data[i].YPos)
		nodePts[i].X = float64(data[i].XPos)
	}

	nodePtsData := nodePts
	p := plot.New()
	p.Add(plotter.NewGrid())
	nodePoints, err := plotter.NewScatter(nodePtsData)
	if err != nil {
		log.Panic(err)
	}

	nodePoints.Shape = draw.CircleGlyph{}
	nodePoints.Color = color.RGBA{G: 255, A: 255}
	nodePoints.Radius = vg.Points(12)

	for index := 0; index < len(endPoints); index++ {
		valA := endPoints[index].ValA
		xA, yA := findXY(valA)
		valB := endPoints[index].ValB
		xB, yB := findXY(valB)
		pts := plotter.XYs{{X: float64(xA), Y: float64(yA)}, {X: float64(xB), Y: float64(yB)}}
		line, err := plotter.NewLine(pts)
		if err != nil {
			log.Panic(err)
		}
		scatter, err := plotter.NewScatter(pts)
		if err != nil {
			log.Panic(err)
		}
		p.Add(line, scatter)
	}
	p.Add(nodePoints)
	for index:=0;index<len(data);index++{
		x:=float64(data[index].XPos) - 0.05
		y:=float64(data[index].YPos) - 0.02
		str:=data[index].Val
		label,err:=plotter.NewLabels(plotter.XYLabels{
			XYs:[]plotter.XY{
				{X:x,Y:y},
			},
			Labels:[]string{str},
		},)
		if err != nil {
			log.Fatalf("Could not create labels plotter: %+v",err)
		}
		p.Add(label)
	}
	// Save the plot to current directory (project folder)
	err = p.Save(1000, 600, path+"tree.png")
	if err != nil {
		log.Panic(err)
	}

	drawGraph(myApp, myWindow, path)
	myWindow.ShowAndRun()
}