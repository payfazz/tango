package container

// ServiceContainer is a struct to handle all client libraries / sdk used in project
type ClientContainer struct {
}

// CreateServiceContainer is a constructor for creating all clients libraries / sdk instance used in the app
func CreateClientContainer() *ClientContainer {
	return &ClientContainer{}
}
