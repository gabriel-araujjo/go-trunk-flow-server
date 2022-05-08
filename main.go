package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gabriel-araujjo/go-trunk-flow-server/pkg/controller"
	"github.com/gabriel-araujjo/go-trunk-flow-server/pkg/features"
	"github.com/gabriel-araujjo/go-trunk-flow-server/pkg/util"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	hello := controller.NewHelloController()

	r.Get("/hello", hello.Hello)
	r.NotFound(util.NotFound)

	f, _ := json.Marshal(features.Features)

	fmt.Printf("{\"server.features\":%s}\n", string(f))

	http.ListenAndServe(":3000", r)

}
