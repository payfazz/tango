package make

import "os"

const (
	MODEL_STUB_FILE      = `/make/v1/template/model.stub`
	REPOSITORY_STUB_FILE = `/make/v1/template/repository.stub`
	PAYLOAD_STUB_FILE    = `/make/v1/template/payload.stub`
	COMMAND_STUB_FILE    = `/make/v1/template/command.stub`
	QUERY_STUB_FILE      = `/make/v1/template/query.stub`
	SERVICE_STUB_FILE    = `/make/v1/template/service.stub`
	DASHBOARD_STUB_FILE  = `/make/v1/template/dashboard.stub`
)

const (
	TYPE_UUID           = `Uuid`
	TYPE_AUTO_INCREMENT = `AutoIncrement`
	TYPE_PLAIN          = `Plain`
)

const (
	DOMAIN_DIR           = "./internal/domain"
	DASHBOARD_DIR        = "./internal/dashboard"
	BACKUP_DIR           = "./internal/old"
	DOMAIN_BACKUP_DIR    = "./internal/old/domain"
	DASHBOARD_BACKUP_DIR = "./internal/old/dashboard"
)

const (
	STRUCTURE_FILE = "./make/structure.yaml"
)

var DIR_FILE_MODE = os.FileMode(0744)
