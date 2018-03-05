package web

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"fmt"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	router := mux.NewRouter()
	for _, route := range routes {
		handlerFunc := log(route.HandlerFunc)
		router.HandleFunc(route.Pattern, handlerFunc).Methods(route.Method)
	}
	var port = getPort()
	if err := http.ListenAndServe(":" + port, router); err != nil {
		panic(err)
	}
}

func log(funcHandler http.HandlerFunc) http.HandlerFunc{
	fmt.Println("Returning the function")
	return func (rw http.ResponseWriter, r *http.Request){
		fmt.Println("New REST request to url: "+r.URL.Path)
		funcHandler(rw, r)
		fmt.Println("Rest request ended")
	}
}

func getPort() string {
	var port = os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}