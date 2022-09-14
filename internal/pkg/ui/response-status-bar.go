package ui

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/pterm/pterm"
	"github.com/rivo/tview"

	"github.com/apicli/apicli/internal/pkg/utils"
)

type ResponseStatusBar struct {
	*tview.Flex
}

func NewResponseStatusBar() *ResponseStatusBar {
	flex := tview.NewFlex()
	flex.SetBackgroundColor(tcell.ColorDefault)
	return &ResponseStatusBar{
		Flex: flex,
	}
}

func (bar *ResponseStatusBar) Reset() {
	bar.SetBackgroundColor(tcell.ColorDefault)
	bar.Clear()
}

func (bar *ResponseStatusBar) ParseResponse(resp *http.Response, contentLength int) {
	bar.Reset()

	var status []string
	status = append(status, "Status: "+utils.GetColoredStatus(resp))
	status = append(status, "Content-Length: "+fmt.Sprint(contentLength))
	status = append(status, "Content-Type: "+resp.Header.Get("Content-Type"))

	bar.AddItem(NewTextView().SetDynamicColors(true).SetText(strings.Join(status, tview.TranslateANSI(pterm.Gray(" | ")))), 0, 1, false)
}

func (bar *ResponseStatusBar) Error(msg string) {
	bar.Reset()
	bar.SetBackgroundColor(tcell.ColorRed)

	textView := NewTextView()
	textView.SetBackgroundColor(tcell.ColorRed)

	bar.AddItem(textView.SetText(msg), 0, 1, false)
}
