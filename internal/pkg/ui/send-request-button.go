package ui

import (
	"errors"
	"io"
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SendRequestButton struct {
	*tview.Button
}

func NewSendRequestButton(methodDropdown *MethodDropDown, urlInput *UrlInput, responseOutput *TextView, container *Flex) *SendRequestButton {
	button := tview.NewButton("Submit")
	button.SetBackgroundColor(tcell.ColorLimeGreen)
	button.SetSelectedFunc(func() {
		_, method := methodDropdown.GetCurrentOption()
		if method == "" {
			CheckErr(errors.New("no method selected"))
			return
		}

		url := urlInput.GetText()
		if url == "" {
			CheckErr(errors.New("no url provided"))
			return
		}

		req, err := http.NewRequest(method, url, nil)
		CheckErr(err)

		resp, err := http.DefaultClient.Do(req)
		CheckErr(err)

		body, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		CheckErr(err)

		container.SetSuccessBorder()
		responseOutput.SetText(string(body))
		responseStatusBar.ParseResponse(resp, len(body))
	})

	return &SendRequestButton{
		Button: button,
	}
}
