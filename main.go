package main

import (
	"github.com/SankeProds/jwtservice/pkg/implementation"
	"github.com/SankeProds/jwtservice/pkg/usecases"
	"github.com/gorilla/mux"
)

func main() {

	// for apps and modules to get configuration params
	conf := new(implementation.EnvOrDefaultConf)

	// data repos
	userRedisRepo := implementation.NewUserRedisRepo(conf)
	signingKeyGetter := implementation.NewSigningKeyGetter(conf)
	jWTGenerator := implementation.NewJWTGenerator(signingKeyGetter)

	// use case handlers
	userCasesHandler := usecases.NewUserUsecase(userRedisRepo)
	sessionCasesHandler := usecases.NewSessionUsecase(userRedisRepo, jWTGenerator)

	// create each app handler
	// App handler knows how to call the use case from  the http call
	apps := [...]implementation.App{
		implementation.NewUserApp(userCasesHandler),
		implementation.NewSessionApp(sessionCasesHandler),
	}

	// Router & Routes
	// Small layer the allows to register each app on the server handler
	muxRouter := mux.NewRouter()
	for _, app := range apps {
		app.RegisterHandlers(muxRouter)
	}

	// Http Server, get the routing info from requestRouter
	server := new(implementation.HttpServer)
	server.Init(conf, muxRouter)
	server.Start()
}
