package container

// ServiceContainer handle all service used in project
type ServiceContainer struct {
}

// CreateServiceContainer construct all services available in the app
func CreateServiceContainer(repositories *RepositoryContainer, grpcClients *GrpcClientContainer, clients *ClientContainer) *ServiceContainer {
	return &ServiceContainer{}
}
