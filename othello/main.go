package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"github.com/Jacalz/fyne-board-games/othello/assets"
)

func main() {
	a := app.NewWithID("com.github.jacalz.fyne-board-games.othello")
	a.SetIcon(assets.Icon)
	w := a.NewWindow("othello")

	bs := &boardStatus{turn: 1, w: &w}

	gameGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(8))
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			gameGrid.AddObject(newBoardIcon(bs, x, y))
		}
	}

	w.SetContent(gameGrid)
	w.ShowAndRun()
}
