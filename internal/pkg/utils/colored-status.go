package utils

import (
	"net/http"

	"github.com/pterm/pterm"
	"github.com/rivo/tview"
)

func GetColoredStatus(r *http.Response) string {
	statusCode := r.StatusCode
	status := r.Status

	var s string
	if statusCode >= 200 && statusCode < 300 {
		s = pterm.FgLightGreen.Sprint(status)
	} else if statusCode >= 300 && statusCode < 400 {
		s = pterm.FgLightYellow.Sprint(status)
	} else if statusCode >= 400 && statusCode < 500 {
		s = pterm.FgLightRed.Sprint(status)
	} else if statusCode >= 500 && statusCode < 600 {
		s = pterm.FgLightRed.Sprint(status)
	}

	return tview.TranslateANSI(s)
}
