package main

import (
	"log"

	"github.com/SankeProds/jwtservice/pkg/implementation"
	"github.com/SankeProds/jwtservice/pkg/usecases"
)

func main() {

	log.Printf("Loading Configuration\n")
	conf := new(implementation.EnvOrDefaultConf)

	// repos
	userRedisRepo := implementation.NewUserRedisRepo(conf)

	// use case handler
	UserCases := usecases.NewUserUsecase(userRedisRepo)
	SessionCases := usecases.NewSessionUsecase(userRedisRepo)

	// Router & Routes
	log.Printf("Creating Request Router\n")
	requestRouter := new(implementation.RequestRouter)
	requestRouter.Init()
	requestRouter.AddApp(implementation.NewUserApp(UserCases))
	requestRouter.AddApp(implementation.NewSessionApp(SessionCases))

	// Server
	log.Printf("Starting server...\n")
	server := new(implementation.HttpServer)
	server.Init(conf, requestRouter)
	server.Start()
}
