package make

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const (
	MODEL_STUB_FILE      = `./make/template/model.stub`
	REPOSITORY_STUB_FILE = `./make/template/repository.stub`
	PAYLOAD_STUB_FILE    = `./make/template/payload.stub`
	COMMAND_STUB_FILE    = `./make/template/command.stub`
	QUERY_STUB_FILE      = `./make/template/query.stub`
	SERVICE_STUB_FILE    = `./make/template/service.stub`
)

const (
	TYPE_UUID           = `Uuid`
	TYPE_AUTO_INCREMENT = `AutoIncrement`
	TYPE_PLAIN          = `Plain`
)

var dirFileMode = os.FileMode(0744)

// GenerateStubs generate all required stubs for CRUD
func GenerateStubs(structure *Structure, baseDir string) {
	domain := strings.ToLower(structure.Model)

	// Make root directory
	dir := fmt.Sprintf("%s/%s", baseDir, domain) // ex: internal/domain/author
	err := os.MkdirAll(dir, dirFileMode)
	if nil != err {
		panic(err)
	}

	generateFile(structure, dir, "model", "model", MODEL_STUB_FILE)
	generateFile(structure, dir, "repository", "repository", REPOSITORY_STUB_FILE)

	if structure.Action.IsCommandNeeded() {
		generateFile(structure, dir, "data", "payload", PAYLOAD_STUB_FILE)
		generateFile(structure, dir, "command", "command", COMMAND_STUB_FILE)
	}

	if structure.Action.IsQueryNeeded() {
		generateFile(structure, dir, "query", "query", QUERY_STUB_FILE)
	}

	if structure.Action.IsQueryNeeded() || structure.Action.IsCommandNeeded() {
		generateFile(structure, dir, "", "service", SERVICE_STUB_FILE)
	}
}

func generateFile(structure *Structure, baseDir string, prefix string, fileName string, stubPath string) {
	// Make dir and file
	insideDir := fmt.Sprintf("%s", baseDir) // ex: internal/domain/author
	if "" != insideDir {
		insideDir = fmt.Sprintf("%s/%s", baseDir, prefix) // ex: internal/domain/author/model
	}
	generatedFile := fmt.Sprintf("%s/%s.go", insideDir, fileName) // ex: internal/domain/author/model/model.go

	_, err := os.Stat(generatedFile)
	if !os.IsNotExist(err) {
		fmt.Println("File", generatedFile, "already exists")
		return
	}

	err = os.MkdirAll(insideDir, dirFileMode)
	if nil != err {
		panic(err)
	}

	f, err := os.Create(generatedFile)
	if nil != err {
		panic(err)
	}

	tmpl, err := template.ParseFiles(stubPath)
	if nil != err {
		panic(err)
	}

	err = tmpl.Execute(f, structure)
	if nil != err {
		panic(err)
	}
}
