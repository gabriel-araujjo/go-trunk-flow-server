package controller

import "net/http"

type ByeController interface {
	Bye(http.ResponseWriter, *http.Request)
}

func NewByeController() ByeController {
	return byeControllerImpl{}
}
