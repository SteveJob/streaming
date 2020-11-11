package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"modernday.com/stevejob/streaming/controller"
	"net/http"
)

func main() {
	fmt.Println("streaming media micro service.")
	router := RegisterHandlers()
	_ = http.ListenAndServe("localhost:9000", router)
}

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	//validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		_, _ = io.WriteString(w, "success! hello streaming")
	})

	router.POST("/user", controller.CreateUser)

	router.POST("/user/:id", controller.Login)

	return router
}