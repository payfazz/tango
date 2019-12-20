package make

import (
	"fmt"
	"strings"

	"github.com/payfazz/tango/cli/util"
)

// StructureMap handle overall structure mapping
type StructureMap struct {
	Structures []*DomainStructure `yaml:"structures"`
}

// DomainStructure handle structure.yaml mapping for domains
type DomainStructure struct {
	Domain string            `yaml:"domain"`
	Models []*ModelStructure `yaml:"models"`
}

func (d *DomainStructure) CamelDomain() string {
	return util.SnakeToCamelCase(d.Domain)
}

func (d *DomainStructure) LowerDomain() string {
	return strings.ToLower(d.Domain)
}

func (d *DomainStructure) Generate(baseDir string) {
	GenerateDomainStubs(d, baseDir)

	for _, model := range d.Models {
		modelDir := fmt.Sprintf("%s/%s", baseDir, d.LowerDomain())
		model.Generate(modelDir)
	}
}

// ModelStructure handle structure.yaml mapping for model inside domains
type ModelStructure struct {
	Name       string   `yaml:"name"`
	Table      string   `yaml:"table"`
	Type       string   `yaml:"idType"`
	DataType   string   `yaml:"idDataType"`
	Timestamps bool     `yaml:"timestamps"`
	SoftDelete bool     `yaml:"softDeletes"`
	Fields     []*Field `yaml:"fields"`
	Action     *Action  `yaml:"action"`
}

func (s *ModelStructure) CamelModel() string {
	return util.SnakeToCamelCase(s.Name)
}

func (s *ModelStructure) LowerModel() string {
	return strings.ToLower(s.Name)
}

// Generate generate current structure and all its child
func (s *ModelStructure) Generate(baseDir string) {
	GenerateModelStubs(s, baseDir)
}

// Field handle field inside structure.yaml mapping
type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func (f *Field) PascalName() string {
	return util.SnakeToPascalCase(f.Name)
}

func (f *Field) CamelName() string {
	return util.SnakeToCamelCase(f.Name)
}

// Action handle functionality that will be generated
type Action struct {
	Create bool `yaml:"create"`
	Read   bool `yaml:"read"`
	Find   bool `yaml:"find"`
	Update bool `yaml:"update"`
	Delete bool `yaml:"delete"`
}

// IsCommandNeeded return true if create or update or delete is requested
func (a *Action) IsCommandNeeded() bool {
	return a.Create || a.Update || a.Delete
}

// IsQueryNeeded return true if read or find is requested
func (a *Action) IsQueryNeeded() bool {
	return a.Read || a.Find
}
