package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Flex struct {
	*tview.Flex
}

func NewFlex() *Flex {
	flex := tview.NewFlex()
	flex.SetBackgroundColor(tcell.ColorDefault)
	return &Flex{
		Flex: flex,
	}
}

func NewFlexWithBorder(title string) *Flex {
	flex := tview.NewFlex()
	flex.SetBackgroundColor(tcell.ColorDefault)
	flex.SetBorder(true)
	if title != "" {
		flex.SetTitle(title)
	}
	return &Flex{
		Flex: flex,
	}
}

func (f *Flex) SetErrorBorder() {
	f.Box.SetBorderColor(tcell.ColorRed)
}

func (f *Flex) SetSuccessBorder() {
	f.Box.SetBorderColor(tcell.ColorGreen)
}

func (f *Flex) SetWarningBorder() {
	f.Box.SetBorderColor(tcell.ColorYellow)
}

func (f *Flex) SetInfoBorder() {
	f.Box.SetBorderColor(tcell.ColorBlue)
}

func (f *Flex) SetDefaultBorder() {
	f.Box.SetBorderColor(tcell.ColorDefault)
}
