package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/Jacalz/fyne-board-games/tic-tac-toe/assets"
)

func main() {
	a := app.NewWithID("com.github.jacalz.fyne-board-games.tic-tac-toe")
	a.SetIcon(assets.Icon)
	w := a.NewWindow("tic-tac-toe")

	bs := &boardStatus{turn: 1, w: &w}

	gameGrid := fyne.NewContainerWithLayout(layout.NewGridLayout(3))
	for i := 0; i < 9; i++ {
		gameGrid.AddObject(newBoardIcon(bs, i))
	}

	startButton := widget.NewButtonWithIcon("Reset Game Board", theme.ViewRefreshIcon(), func() {
		bs.Cleanup()

		for i := 0; i < 9; i++ {
			gameGrid.Objects[i] = newBoardIcon(bs, i)
		}
	})

	w.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), startButton, gameGrid))
	w.ShowAndRun()
}
