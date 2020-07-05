package configs

import (
	"database/sql"
	"fmt"
	"gomart-api/utils"
	"log"
)

var (
	DbEngine     string
	DbUser       string
	DbPass       string
	DbHost       string
	DbPort       string
	DbConnection string
	DbSchema     string
)

//ConnectDB to connect into database
func ConnectDB() (*sql.DB, error) {

	DbEngine = utils.GetCustomConf("DbEngine", "mysql")
	DbUser = utils.GetCustomConf("DbUser", "root")
	DbPass = utils.GetCustomConf("DbPass", "P@ssw0rd")
	DbHost = utils.GetCustomConf("DbHost", "localhost")
	DbPort = utils.GetCustomConf("DbPort", "8080")
	DbConnection = utils.GetCustomConf("DbConnection", "@tcp")
	DbSchema = utils.GetCustomConf("DbSchema", "enigma_school_api")

	dataSource := fmt.Sprintf("%s:%s%s(%s:%s)/%s", DbUser, DbPass, DbConnection, DbHost, DbPort, DbSchema)

	DB, _ := sql.Open(DbEngine, dataSource)

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	return DB, nil
}
