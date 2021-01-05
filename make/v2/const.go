package make

import "os"

const (
	BASE_STUB_FILE_GLOB   = `/make/v2/template/*.stub`
	METHOD_STUB_FILE_GLOB = `/make/v2/template/methods/*.stub`

	DOMAIN_DIR = "./internal/domain"
)

var DIR_FILE_MODE = os.FileMode(0744)
