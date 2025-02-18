package main

import (
	"notification-service/cmd"
	"notification-service/helpers"
)

func main() {

	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupPostgreSQL()

	// run grpc
	go cmd.ServeGRPC()

	// run http
	// cmd.ServeHTTP()
}
