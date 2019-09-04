package make

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/payfazz/go-apt/pkg/fazzcommon/formatter"
)

const (
	MODEL_STUB_FILE                = `./make/template/model.stub`
	REPOSITORY_STUB_FILE           = `./make/template/repository.stub`
	REPOSITORY_INTERFACE_STUB_FILE = `./make/template/repositoryinterface.stub`
	PAYLOAD_STUB_FILE              = `./make/template/payload.stub`
	COMMAND_STUB_FILE              = `./make/template/command.stub`
	QUERY_STUB_FILE                = `./make/template/query.stub`
	SERVICE_STUB_FILE              = `./make/template/service.stub`
)

var dirFileMode = os.FileMode(0755)
var goFileMode = os.FileMode(0644)

// GenerateStubs generate all required stubs for CRUD
func GenerateStubs(structure *Structure) {
	domain := strings.ToLower(structure.Model)

	oldString, newString := generateStructureReplacements(structure, domain)

	// Make root directory
	dir := fmt.Sprintf("./internal/domain/%s", domain)
	err := os.MkdirAll(dir, dirFileMode)
	if nil != err {
		panic(err)
	}

	baseMeta := &meta{
		Directory: dir,
		Old:       oldString,
		New:       newString,
	}

	baseMeta.generateFile("model", "model", MODEL_STUB_FILE)
	baseMeta.generateFile("repository", "interface", REPOSITORY_INTERFACE_STUB_FILE)
	baseMeta.generateFile("repository", "postgres", REPOSITORY_STUB_FILE)
	baseMeta.generateFile("data", "payload", PAYLOAD_STUB_FILE)
	baseMeta.generateFile("command", "command", COMMAND_STUB_FILE)
	baseMeta.generateFile("query", "query", QUERY_STUB_FILE)
	baseMeta.generateFile("", "service", SERVICE_STUB_FILE)
}

type meta struct {
	Directory string
	Old       []string
	New       []string
}

func (m *meta) generateFile(dirPrefix string, fileName string, stubPath string) {
	// Make dir and file
	insideDir := dirPrefix
	if "" != insideDir {
		insideDir = fmt.Sprintf("/%s", insideDir)
	}
	dir := fmt.Sprintf("%s%s", m.Directory, insideDir)
	generatedFile := fmt.Sprintf("%s/%s.go", dir, fileName)

	_, err := os.Stat(generatedFile)
	if !os.IsNotExist(err) {
		fmt.Println("File", generatedFile, "already exists")
		return
	}

	err = os.MkdirAll(dir, dirFileMode)
	if nil != err {
		panic(err)
	}

	content, err := ioutil.ReadFile(stubPath)
	if nil != err {
		panic(err)
	}
	stub := string(content)

	// Replace stub
	stub = formatter.ReplaceStrings(stub, m.Old, m.New)
	err = ioutil.WriteFile(generatedFile, []byte(stub), goFileMode)
	if nil != err {
		panic(err)
	}
}
