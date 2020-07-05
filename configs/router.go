package configs

import (
	"fmt"
	"gomart-api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router) {
	ServerHost := utils.GetCustomConf("ServerHost", "localhost")
	ServerPort := utils.GetCustomConf("ServerPort", "3500")

	serverSource := fmt.Sprintf("%s:%s", ServerHost, ServerPort)

	fmt.Println("Starting Web Server At Port : " + ServerPort)
	err := http.ListenAndServe(serverSource, router)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
