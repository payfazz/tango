package container

// AppContainer is a struct to handle all requirement for app to run properly
type AppContainer struct {
	Clients     *ClientContainer
	Middlewares *MiddlewareContainer
	Services    *ServiceContainer
}

// CreateServiceContainer is a constructor for creating all requirement for app
func CreateAppContainer() *AppContainer {
	repositories := CreateRepositoryContainer()
	clients := CreateClientContainer()

	return &AppContainer{
		Clients:     clients,
		Middlewares: CreateMiddlewareContainer(),
		Services:    CreateServiceContainer(repositories, clients),
	}
}
