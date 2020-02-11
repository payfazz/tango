package container

import "github.com/payfazz/tango/transport/http/container"

// AppContainer handle all requirement for app to run properly
type AppContainer struct {
	Clients        *ClientContainer
	GrpcClients    *GrpcClientContainer
	HttpMiddleware *container.MiddlewareContainer
	Services       *ServiceContainer
}

// CreateServiceContainer construct all requirement for app
func CreateAppContainer() *AppContainer {
	repositories := CreateRepositoryContainer()
	clients := CreateClientContainer()
	grpcClients := CreateGrpcClientContainer()

	return &AppContainer{
		Clients:        clients,
		GrpcClients:    grpcClients,
		HttpMiddleware: container.CreateMiddlewareContainer(),
		Services:       CreateServiceContainer(repositories, grpcClients, clients),
	}
}
