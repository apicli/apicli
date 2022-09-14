package ui

import "gitlab.com/tslocum/cview"

func GetMainView() *Flex {
	wrapper := NewFlex()
	leftSidebar := NewFlexWithBorder("Saved Requests")
	requestBar := NewFlexWithBorder("")
	requestSettings := NewFlexWithBorder("Request Parameters")
	responseContainer := NewFlexWithBorder("Response")
	rightSidebar := NewFlexWithBorder("")
	// statusBar := NewFlex()

	leftContainer := NewFlex()
	middleContainer := NewFlex()
	rightContainer := NewFlex()

	wrapper.SetFullScreen(true)
	wrapper.SetDirection(cview.FlexColumn)

	wrapper.AddItem(leftContainer, 0, 1, false)
	wrapper.AddItem(middleContainer, 0, 3, false)
	wrapper.AddItem(rightContainer, 20, 0, false)

	leftContainer.SetDirection(cview.FlexRow)
	leftContainer.AddItem(leftSidebar, 0, 1, false)

	middleContainer.SetDirection(cview.FlexRow)
	middleContainer.AddItem(requestBar, 3, 0, false)
	middleContainer.AddItem(requestSettings, 0, 2, false)
	middleContainer.AddItem(responseContainer, 0, 3, false)

	rightContainer.SetDirection(cview.FlexRow)
	rightContainer.AddItem(rightSidebar, 0, 1, false)

	return wrapper
}
