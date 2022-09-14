package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Pages struct {
	*tview.Pages
}

func NewPages() *Pages {
	pages := tview.NewPages()
	pages.SetBackgroundColor(tcell.ColorDefault)
	return &Pages{
		Pages: pages,
	}
}
