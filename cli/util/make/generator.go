package make

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

	QUERY_TAG   = `qu-tag`
	COMMAND_TAG = `co-tag`
	READ_TAG    = `re-tag`
	FIND_TAG    = `fi-tag`
	CREATE_TAG  = `cr-tag`
	UPDATE_TAG  = `up-tag`
	DELETE_TAG  = `de-tag`
)

var dirFileMode = os.FileMode(0755)
var goFileMode = os.FileMode(0644)

// GenerateStubs generate all required stubs for CRUD
func GenerateStubs(structure *Structure, baseDir string) {
	domain := strings.ToLower(structure.Model)

	oldString, newString := generateStructureReplacements(structure, domain)

	// Make root directory
	dir := fmt.Sprintf("%s/%s", baseDir, domain)
	err := os.MkdirAll(dir, dirFileMode)
	if nil != err {
		panic(err)
	}

	baseMeta := &meta{
		directory: dir,
		old:       oldString,
		new:       newString,
		unused:    []string{},
	}

	baseMeta.parseAction(structure.Action)

	baseMeta.generateFile("model", "model", MODEL_STUB_FILE)
	baseMeta.generateFile("repository", "interface", REPOSITORY_INTERFACE_STUB_FILE)
	baseMeta.generateFile("repository", "postgres", REPOSITORY_STUB_FILE)

	if structure.Action.IsCommandNeeded() {
		baseMeta.generateFile("data", "payload", PAYLOAD_STUB_FILE)
		baseMeta.generateFile("command", "command", COMMAND_STUB_FILE)
	}

	if structure.Action.IsQueryNeeded() {
		baseMeta.generateFile("query", "query", QUERY_STUB_FILE)
	}

	if structure.Action.IsQueryNeeded() || structure.Action.IsCommandNeeded() {
		baseMeta.generateFile("", "service", SERVICE_STUB_FILE)
	}
}

type meta struct {
	directory string
	old       []string
	new       []string
	unused    []string
}

func (m *meta) parseAction(action *Action) {
	if action.Read {
		m.old = append(m.old, fmt.Sprintf("{{%s}}", READ_TAG), fmt.Sprintf("{{end-%s}}", READ_TAG))
		m.new = append(m.new, "", "")
	} else {
		m.unused = append(m.unused, READ_TAG)
	}

	if action.Find {
		m.old = append(m.old, fmt.Sprintf("{{%s}}", FIND_TAG), fmt.Sprintf("{{end-%s}}", FIND_TAG))
		m.new = append(m.new, "", "")
	} else {
		m.unused = append(m.unused, FIND_TAG)
	}

	if action.Create {
		m.old = append(m.old, fmt.Sprintf("{{%s}}", CREATE_TAG), fmt.Sprintf("{{end-%s}}", CREATE_TAG))
		m.new = append(m.new, "", "")
	} else {
		m.unused = append(m.unused, CREATE_TAG)
	}

	if action.Update {
		m.old = append(m.old, fmt.Sprintf("{{%s}}", UPDATE_TAG), fmt.Sprintf("{{end-%s}}", UPDATE_TAG))
		m.new = append(m.new, "", "")
	} else {
		m.unused = append(m.unused, UPDATE_TAG)
	}

	if action.Delete {
		m.old = append(m.old, fmt.Sprintf("{{%s}}", DELETE_TAG), fmt.Sprintf("{{end-%s}}", DELETE_TAG))
		m.new = append(m.new, "", "")
	} else {
		m.unused = append(m.unused, DELETE_TAG)
	}

	if action.IsCommandNeeded() {
		m.old = append(m.old, fmt.Sprintf("{{%s}}", COMMAND_TAG), fmt.Sprintf("{{end-%s}}", COMMAND_TAG))
		m.new = append(m.new, "", "")
	} else {
		m.unused = append(m.unused, COMMAND_TAG)
	}

	if action.IsQueryNeeded() {
		m.old = append(m.old, fmt.Sprintf("{{%s}}", QUERY_TAG), fmt.Sprintf("{{end-%s}}", QUERY_TAG))
		m.new = append(m.new, "", "")
	} else {
		m.unused = append(m.unused, QUERY_TAG)
	}
}

func (m *meta) generateFile(dirPrefix string, fileName string, stubPath string) {
	// Make dir and file
	insideDir := dirPrefix
	if "" != insideDir {
		insideDir = fmt.Sprintf("/%s", insideDir)
	}
	dir := fmt.Sprintf("%s%s", m.directory, insideDir)
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

	// Remove unused content and tags
	for _, v := range m.unused {
		regex := fmt.Sprintf(`{{%s}}(.|\s)*?{{end-%s}}`, v, v)
		re := regexp.MustCompile(regex)
		stub = re.ReplaceAllString(stub, "")
	}

	// Replace stub
	stub = formatter.ReplaceStrings(stub, m.old, m.new)
	err = ioutil.WriteFile(generatedFile, []byte(stub), goFileMode)
	if nil != err {
		panic(err)
	}
}
