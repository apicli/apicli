package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StatusBar struct {
	*tview.Box
}

func NewStatusBar() *StatusBar {
	return &StatusBar{
		Box: tview.NewBox().SetBackgroundColor(tcell.ColorLightCyan),
	}
}
