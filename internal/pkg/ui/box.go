package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Box struct {
	*tview.Box
}

func NewBox() *Box {
	return &Box{
		Box: tview.NewBox().
			SetBackgroundColor(tcell.ColorDefault).
			SetBorder(true),
	}
}
