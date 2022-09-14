package ui

type MethodDropDown struct {
	*DropDown
}

func NewMethodDropDown(method string) *MethodDropDown {
	methodDropDown := &MethodDropDown{
		DropDown: NewDropDown(),
	}
	httpMethods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE", "CONNECT"}
	methodDropDown.DropDown.SetOptions(httpMethods, nil)
	methodDropDown.SetLabel("Method: ")

	var methodId int
	for i, m := range httpMethods {
		if m == method {
			methodId = i
		}
	}
	methodDropDown.SetCurrentOption(methodId)

	return methodDropDown
}
