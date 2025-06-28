package main

import (
	"perpus-app/cmd"
	"perpus-app/helpers"
)

func main() {

	// load setup
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load database
	// helpers.SetupMySQL()

	//run HTTP
	cmd.ServeHTTP()
}
