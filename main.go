package main

import (
	"alfdwirhmn/bioskop/cmd"
	"alfdwirhmn/bioskop/internal/wire"
	"alfdwirhmn/bioskop/pkg/database"
	"alfdwirhmn/bioskop/pkg/utils"
	"log"
)

func main() {
	// read config
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	// init logger
	logger, err := utils.InitLogger(config.PathLogg, config)
	if err != nil {
		log.Fatal(err)
	}

	// init database
	db, err := database.InitDB(config.DB)
	if err != nil {
		logger.Fatal("failed to connect database")
	}

	router := wire.SetupRouter(db, logger, config)

	// run server
	cmd.ApiServer(config, logger, router)
}
