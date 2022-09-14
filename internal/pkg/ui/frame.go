package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Frame struct {
	*tview.Frame
}

func NewFrame() *Frame {
	frame := tview.NewFrame(nil)
	frame.SetBackgroundColor(tcell.ColorDefault)
	return &Frame{
		Frame: frame,
	}
}
