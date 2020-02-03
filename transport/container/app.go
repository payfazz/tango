package container

import "github.com/payfazz/tango/transport/http/container"

// AppContainer is a struct to handle all requirement for app to run properly
type AppContainer struct {
	Clients     *ClientContainer
	GrpcClients *GrpcClientContainer
	Middlewares *container.MiddlewareContainer
	Services    *ServiceContainer
}

// CreateServiceContainer is a constructor for creating all requirement for app
func CreateAppContainer() *AppContainer {
	repositories := CreateRepositoryContainer()
	clients := CreateClientContainer()
	grpcClients := CreateGrpcClientContainer()

	return &AppContainer{
		Clients:     clients,
		GrpcClients: grpcClients,
		Middlewares: container.CreateMiddlewareContainer(),
		Services:    CreateServiceContainer(repositories, grpcClients, clients),
	}
}
