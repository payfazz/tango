package container

// RepositoryContainer handle all repository used in project
type RepositoryContainer struct {
}

// CreateRepositoryContainer construct all repositories available in the app
func CreateRepositoryContainer() *RepositoryContainer {
	return &RepositoryContainer{}
}
