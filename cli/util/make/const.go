package make

import "os"

const (
	MODEL_STUB_FILE      = `./make/template/model.stub`
	REPOSITORY_STUB_FILE = `./make/template/repository.stub`
	PAYLOAD_STUB_FILE    = `./make/template/payload.stub`
	COMMAND_STUB_FILE    = `./make/template/command.stub`
	QUERY_STUB_FILE      = `./make/template/query.stub`
	SERVICE_STUB_FILE    = `./make/template/service.stub`
	DASHBOARD_STUB_FILE  = `./make/template/dashboard.stub`
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
	STRUCTURE_PATH = "./make/structure.yaml"
)

var DIR_FILE_MODE = os.FileMode(0744)
