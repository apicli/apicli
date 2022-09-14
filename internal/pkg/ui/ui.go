package ui

import (
	"github.com/rivo/tview"
)

var app *tview.Application

var (
	ResponseOutput    = NewTextView()
	Views             = NewPages()
	MainContent       *Flex
	responseStatusBar = NewResponseStatusBar()
)

func Start() error {
	app = tview.NewApplication()

	sidebar := NewFlex().SetBorder(true).SetTitle("Left (1/3 x width of Top)")

	methodDropDown := NewMethodDropDown("GET")
	urlInput := NewUrlInput()
	urlInput.SetText("https://example.com") // TODO: REMOVE
	response := NewFlexWithBorder("Response")

	submitButton := NewSendRequestButton(methodDropDown, urlInput, ResponseOutput, response)

	requestUrlForm := NewFlex().SetDirection(tview.FlexRow).
		AddItem(NewFlexWithBorder("").SetDirection(tview.FlexColumn).
			AddItem(methodDropDown, 17, 0, false).
			AddItem(urlInput, 0, 1, false).
			AddItem(submitButton, 10, 0, false), 3, 0, false)

	MainContent = &Flex{
		Flex: NewFlex().
			AddItem(sidebar, 0, 2, false).
			AddItem(requestUrlForm.
				AddItem(NewBox().SetTitle("Request Params"), 0, 2, false).
				AddItem(response.
					AddItem(NewFlex().SetDirection(tview.FlexRow).
						AddItem(responseStatusBar, 1, 0, false). // Response status
						AddItem(ResponseOutput, 0, 1, false), 0, 1, false), 0, 3, false), 0, 5, false).
			AddItem(NewBox().SetTitle("Right (20 cols)"), 20, 1, false),
	}

	mainPage := NewFlex().SetDirection(tview.FlexRow).
		AddItem(MainContent, 0, 1, false).
		AddItem(NewStatusBar(), 1, 0, false)

	Views.AddPage("main", mainPage, true, true)

	if err := app.SetRoot(Views, true).EnableMouse(true).Run(); err != nil {
		return err
	}

	return nil
}

func Stop() {
	app.Stop()
}
