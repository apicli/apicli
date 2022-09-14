package ui

import (
	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

type Flex struct {
	*cview.Flex
}

func NewFlex() *Flex {
	flex := cview.NewFlex()
	flex.SetBackgroundColor(tcell.ColorDefault)
	return &Flex{
		Flex: flex,
	}
}

func NewFlexWithBorder(title string) *Flex {
	flex := NewFlex()
	flex.SetBorder(true)
	flex.SetBorderColor(tcell.ColorWhite)
	flex.SetBorderColorFocused(tcell.ColorRed)
	flex.ShowFocus(true)
	if title != "" {
		flex.SetTitle(title)
	}

	return flex
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
