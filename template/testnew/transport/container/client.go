package container

// ClientContainer handle all clients libraries / sdk used in project
type ClientContainer struct {
}

// CreateClientContainer construct all clients libraries / sdk instance used in the app
func CreateClientContainer() *ClientContainer {
	return &ClientContainer{}
}
