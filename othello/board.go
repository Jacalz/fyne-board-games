package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/Jacalz/fyne-board-games/othello/assets"
)

var emptyBoard = [8][8]bool{}

type boardStatus struct {
	BoardPressed   [8][8]bool
	Player1Markers int
	Player2Markers int
	turn           int
	finished       bool

	w *fyne.Window
}

func (bs *boardStatus) Cleanup() {
	bs.finished = false
	bs.Player1Markers = 0
	bs.Player2Markers = 0
	bs.BoardPressed = emptyBoard
	bs.turn = 1
}

type boardIcon struct {
	widget.Icon
	x, y   int
	status *boardStatus
}

func (b *boardIcon) Tapped(ev *fyne.PointEvent) {
	if b.status.BoardPressed[b.x][b.y] || b.status.finished {
		return
	} else if b.status.turn < 3 {
		// TODO: Only support laying in the middle at first.
	}

	if b.status.turn%2 == 0 {
		b.SetResource(assets.MarkerBlack)
		b.status.Player2Markers++
	} else {
		b.SetResource(assets.MarkerWhite)
		b.status.Player1Markers++
	}

	b.status.turn++
}

func (b *boardIcon) MinSize() fyne.Size {
	return fyne.NewSize(64, 64)
}

func newBoardIcon(bs *boardStatus, x, y int) *boardIcon {
	b := &boardIcon{status: bs, x: x, y: y}
	b.SetResource(assets.Board)
	b.ExtendBaseWidget(b)
	return b
}
