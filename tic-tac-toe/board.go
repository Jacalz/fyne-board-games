package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/Jacalz/fyne-board-games/tic-tac-toe/assets"
)

var emptyBoard = [9]bool{false, false, false, false, false, false, false, false, false}

type boardStatus struct {
	BoardPlayer1 [9]bool
	BoardPlayer2 [9]bool
	BoardPressed [9]bool
	turn         int
	finished     bool

	w *fyne.Window
}

func (bs *boardStatus) CheckResult(player [9]bool) bool {
	// Switch statement with all possible combinations for winning.
	switch {
	case player[0] && player[1] && player[2]:
		return true
	case player[3] && player[4] && player[5]:
		return true
	case player[6] && player[7] && player[8]:
		return true
	case player[0] && player[3] && player[6]:
		return true
	case player[1] && player[4] && player[7]:
		return true
	case player[2] && player[5] && player[8]:
		return true
	case player[0] && player[4] && player[8]:
		return true
	case player[2] && player[4] && player[6]:
		return true
	}

	return false
}

func (bs *boardStatus) NewBoardClick(i int) {
	if bs.turn%2 == 0 {
		bs.BoardPlayer1[i] = true
	} else {
		bs.BoardPlayer2[i] = true
	}

	bs.BoardPressed[i] = true

	if bs.turn >= 5 && bs.turn%2 == 0 {
		if bs.CheckResult(bs.BoardPlayer1) {
			dialog.ShowInformation("Player 1 has won!", "Congratulations to player 1 for winning.", *bs.w)
			bs.finished = true
		}
	} else if bs.turn >= 5 {
		if bs.CheckResult(bs.BoardPlayer2) {
			dialog.ShowInformation("Player 2 has won!", "Congratulations to player 2 for winning.", *bs.w)
			bs.finished = true
		}
	}
}

func (bs *boardStatus) Cleanup() {
	bs.finished = false
	bs.BoardPlayer1 = emptyBoard
	bs.BoardPlayer2 = emptyBoard
	bs.BoardPressed = emptyBoard
	bs.turn = 1
}

type boardIcon struct {
	widget.Icon
	index  int
	status *boardStatus
}

func (b *boardIcon) Tapped(ev *fyne.PointEvent) {
	if b.status.BoardPressed[b.index] || b.status.finished {
		return
	}

	b.status.NewBoardClick(b.index)

	if b.status.turn%2 == 0 {
		b.SetResource(assets.Cross)
	} else {
		b.SetResource(assets.Circle)
	}

	b.status.turn++
}

func (b *boardIcon) MinSize() fyne.Size {
	return fyne.NewSize(128, 128)
}

func newBoardIcon(bs *boardStatus, i int) *boardIcon {
	b := &boardIcon{status: bs, index: i}
	b.SetResource(theme.ViewFullScreenIcon())
	b.ExtendBaseWidget(b)
	return b
}
