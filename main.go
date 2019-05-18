package main

import (
	"github.com/SankeProds/jwtservice/pkg/implementation"
	"github.com/SankeProds/jwtservice/pkg/interfaces"
	"github.com/SankeProds/jwtservice/pkg/usecases"
)

func main() {

	// data repos and interfaces
	conf := new(implementation.EnvOrDefaultConf)

	authUserPostgresStorage := implementation.NewAuthUserPostgresStorage(conf)
	signingMethod := implementation.NewJWTGenerator(conf)

	userStorage := interfaces.NewUserStorage(authUserPostgresStorage)
	tokenGenerator := interfaces.NewTokenGenerator(signingMethod)

	// use case handlers
	registerUserUCHandler := usecases.NewRegisterUserUC(userStorage)
	loginUCHandler := usecases.NewLoginUC(userStorage, tokenGenerator)

	// create each app handler
	// App handler knows how to call the use case from  the http call
	apps := [...]implementation.HttpApp{
		implementation.NewRegisterUserApp(registerUserUCHandler),
		implementation.NewLoginApp(loginUCHandler),
	}

	httpRouter := implementation.NewHttpRouter()
	for _, app := range apps {
		httpRouter.RegisterApp(app)
	}

	// Http Server, get the routing info from requestRouter
	server := new(implementation.HttpServer)
	server.Init(conf, httpRouter)
	server.Start()
}
