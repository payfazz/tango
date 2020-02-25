package make

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// GenerateDomainStubs generate all required stubs for domain service
func GenerateDomainStubs(structure *DomainStructure, baseDir string) {
	// Make root directory
	dir := fmt.Sprintf("%s/%s", baseDir, strings.ToLower(structure.Domain)) // ex: internal/domain/author
	err := os.MkdirAll(dir, DIR_FILE_MODE)
	if nil != err {
		panic(err)
	}

	generateFile(structure, dir, "service", SERVICE_STUB_FILE)
}

// GenerateModelStubs generate all required stubs for CRUD
func GenerateModelStubs(structure *ModelStructure, baseDir string) {
	modelName := strings.ToLower(structure.Name)

	// Make root directory
	dir := fmt.Sprintf("%s/%s", baseDir, modelName) // ex: internal/domain/inventory/author
	err := os.MkdirAll(dir, DIR_FILE_MODE)
	if nil != err {
		panic(err)
	}

	generateFile(structure, dir, "model", MODEL_STUB_FILE)
	generateFile(structure, dir, "repository", REPOSITORY_STUB_FILE)

	if structure.Action.IsCommandNeeded() {
		generateFile(structure, dir, "payload", PAYLOAD_STUB_FILE)
		generateFile(structure, dir, "command", COMMAND_STUB_FILE)
	}

	if structure.Action.IsQueryNeeded() {
		generateFile(structure, dir, "query", QUERY_STUB_FILE)
	}

	dashboardDir := fmt.Sprintf("%s/%s", DASHBOARD_DIR, strings.ToLower(structure.Name))
	if structure.Action.IsCommandNeeded() || structure.Action.IsQueryNeeded() {
		generateFile(structure, dashboardDir, "dashboard", DASHBOARD_STUB_FILE)
	}
}

func generateFile(structure interface{}, baseDir string, fileName string, stubPath string) {
	// Make dir and file
	generatedFile := fmt.Sprintf("%s/%s.go", baseDir, fileName) // ex: internal/domain/inventory/author/model.go
	fmt.Println("creating file:", generatedFile)

	_, err := os.Stat(generatedFile)
	if !os.IsNotExist(err) {
		fmt.Println("File", generatedFile, "already exists")
		return
	}

	err = os.MkdirAll(baseDir, DIR_FILE_MODE)
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
