package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type DropDown struct {
	*tview.DropDown
}

func NewDropDown() *DropDown {
	dropdown := tview.NewDropDown()
	dropdown.SetBackgroundColor(tcell.ColorDefault)

	return &DropDown{
		DropDown: dropdown,
	}
}
