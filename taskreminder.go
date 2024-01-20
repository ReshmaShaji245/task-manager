package main

import (
	"log"
	"taskreminder/config"
	postgreservice "taskreminder/database/postgres"
	taskhandler "taskreminder/handler/taskhandler"
	"taskreminder/server"
	"taskreminder/services/taskchecksvc"
	tasksvc "taskreminder/services/tasksvc"
	"taskreminder/utils"
)

var (
	serverConfig *config.ServerConfig
	PostgresSvc  *postgreservice.PostgresDBService
	logger       *utils.Logger
)

func init() {

	log.Println("initiating task management server")

	// parse conf
	log.Println("Parsing Configurations")

	var err error

	serverConfig, err = config.GetServerConfig()
	if err != nil {
		log.Panicln("EROR : ", err)
	}

	// initiate logger
	logger = utils.InitiateLoggers(serverConfig.LogFilePath)

	// initiate db
	PostgresSvc = postgreservice.NewPostgresDBService(&serverConfig.DbConfig)
}

func main() {

	logger.Info.Println("starting server")

	// initiate services and their handlers
	taskSvc := tasksvc.NewTaskService(PostgresSvc, logger)
	taskHandler := taskhandler.NewTaskHandler(taskSvc, logger)
	taskCheckSvc := taskchecksvc.NewTaskCheckSvc(PostgresSvc, logger)

	// initiate server
	handlers := []*server.ServerHandlerMap{
		server.NewServerHandlerMap("/v1/task", taskHandler),
	}
	chiApp := server.NewServer()

	server := &server.Server{
		Port:        *serverConfig.Port,
		ChiApp:      chiApp,
		APIRootPath: "/api",
		Handlers:    handlers,
	}

	waitforShutdownInterrupt := server.StartServer()
	logger.Info.Println("Server started")
	logger.Info.Println("Server is running on:", server.Port)
	taskCheckSvc.Start()

	select {
	case <-waitforShutdownInterrupt:
		logger.Info.Println("Server Stopped")
		goto stop
	}

stop:

	logger.Info.Println("initiating server stop sequence")

	server.ShutdownGracefully()

	logger.Info.Println("Server Stopped")
}
