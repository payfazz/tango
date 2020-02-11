package container

// GrpcClientContainer handle all grpc clients libraries / sdk used in project
type GrpcClientContainer struct {
}

// CreateGrpcClientContainer construct all grpc clients libraries / sdk used in project
func CreateGrpcClientContainer() *GrpcClientContainer {
	return &GrpcClientContainer{}
}
