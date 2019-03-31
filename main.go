package main

import (
	"log"

	"github.com/SankeProds/jwtservice/pkg/implementation"
	"github.com/SankeProds/jwtservice/pkg/usecases"
)

func main() {

	// for apps and modules to get configuration params
	log.Printf("Configuration")
	conf := new(implementation.EnvOrDefaultConf)

	// data repos
	log.Printf("Configuration")
	userRedisRepo := implementation.NewUserRedisRepo(conf)

	// use case handlers
	userCasesHandler := usecases.NewUserUsecase(userRedisRepo)
	sessionCasesHandler := usecases.NewSessionUsecase(userRedisRepo)

	// create each app handler
	// App handler knows how to call the use case from  the http call
	apps := [...]implementation.App{
		implementation.NewUserApp(userCasesHandler),
		implementation.NewSessionApp(sessionCasesHandler),
	}

	// Router & Routes
	// Small layer the allows to register each app on the server handler
	log.Printf("Creating Request Router\n")
	requestRouter := new(implementation.RequestRouter)
	requestRouter.Init()
	for _, app := range apps {
		requestRouter.AddApp(app)
	}

	// Http Server, get the routing info from requestRouter
	log.Printf("Starting server...\n")
	server := new(implementation.HttpServer)
	server.Init(conf, requestRouter)
	server.Start()
}
