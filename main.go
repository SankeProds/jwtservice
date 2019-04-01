package main

import (
	"github.com/SankeProds/jwtservice/pkg/implementation"
	"github.com/SankeProds/jwtservice/pkg/interfaces"
	"github.com/SankeProds/jwtservice/pkg/usecases"
)

func main() {

	// for apps and modules to get configuration params
	conf := new(interfaces.EnvOrDefaultConf)

	redisStringStorage := implementation.NewRedisStringStorage(conf)

	// data repos
	userStorage := interfaces.NewUserStorage(redisStringStorage)

	signingKeyGetter := interfaces.NewSigningKeyGetter(conf)
	jWTGenerator := interfaces.NewJWTGenerator(signingKeyGetter)

	// use case handlers
	userCasesHandler := usecases.NewUserUsecase(userStorage)
	sessionCasesHandler := usecases.NewSessionUsecase(userStorage, jWTGenerator)

	httpRouter := implementation.NewHttpRouter()

	// create each app handler
	// App handler knows how to call the use case from  the http call
	apps := [...]implementation.HttpApp{
		implementation.NewUserApp(userCasesHandler),
		implementation.NewSessionApp(sessionCasesHandler),
	}

	for _, app := range apps {
		httpRouter.RegisterApp(app)
	}

	// Http Server, get the routing info from requestRouter
	server := new(implementation.HttpServer)
	server.Init(conf, httpRouter)
	server.Start()
}
