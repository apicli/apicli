package ui

func ShowDialog(msg string) {

	//modal := tview.NewModal().SetText(msg).AddButtons([]string{"OK"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
	//	app.SetRoot(MainContent, true)
	//})
	//app.SetRoot(modal, true)
}

func ShowErrorDialog(msg string) {
	ShowDialog(msg)
}

func CheckErr(err error) {
	if err == nil {
		return
	}

	responseStatusBar.Error(err.Error())
}
