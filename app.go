package main

import (
	"gomart-api/configs"
	"gomart-api/master"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, _ := configs.ConnectDB()
	router := configs.CreateRouter()
	master.Init(db, router)
	configs.RunServer(router)
}
