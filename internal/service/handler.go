package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

type Context struct {
}

func Error(rw web.ResponseWriter, req *web.Request, err interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered panic:", err)
			return
		}
		fmt.Println("no panic recovered")
	}()
}

func (c *Context) HandleRoot(rw web.ResponseWriter, req *web.Request) {

	res := map[string]string{
		"message": "Hello World",
	}

	resJson, err := json.Marshal(res)

	// handle error
	if err != nil {
		JSONHandleError(rw, appError{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// set headers for json content
	rw.Header().Set("Content-Type", "application/json")

	// write json data
	rw.Write(resJson)
}

func initRouter(basePath string) *web.Router {
	// create router
	router := web.New(Context{})

	// ovveride gocraft defualt error handler
	router.Error(Error)

	// add middlewares
	router.Middleware(web.LoggerMiddleware)
	// router.Middleware(web.ShowErrorsMiddleware)

	// handle routes
	router.Get("/", (*Context).HandleRoot)

	return router
}
