package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TextView struct {
	*tview.TextView
}

func NewTextView() *TextView {
	textView := tview.NewTextView()
	textView.SetBackgroundColor(tcell.ColorDefault)
	return &TextView{
		TextView: textView,
	}
}
