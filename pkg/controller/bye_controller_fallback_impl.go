//go:build !(controller.bye || dev)

package controller

import (
	"net/http"

	"github.com/gabriel-araujjo/go-trunk-flow-server/pkg/util"
)

type byeControllerImpl struct{}

func (byeControllerImpl) Bye(w http.ResponseWriter, r *http.Request) {
	util.NotFound(w, r)
}
