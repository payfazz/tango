package container

// RepositoryContainer is a struct to handle all repository used in project
type RepositoryContainer struct {
}

// CreateRepositoryContainer is a constructor for creating all repositories available in the app
func CreateRepositoryContainer() *RepositoryContainer {
	return &RepositoryContainer{}
}
