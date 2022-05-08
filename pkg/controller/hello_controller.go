package controller

import (
	"net/http"
)

type HelloController interface {
	Hello(http.ResponseWriter, *http.Request)
}

func NewHelloController() HelloController {
	return helloControllerImpl{}
}
