package web

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"goworkshop/model"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	http.HandleFunc("/books", httpHandler)
	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		if data, err := json.Marshal(model.Books); err != nil {
			fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
		} else {
			fmt.Fprintln(w, string(data))
		}
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "{\"message\":\"Method not supported!\"}")
		break
	}
}

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}
