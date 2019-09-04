package make

import (
	"fmt"

	"github.com/payfazz/tango/cli/util"
)

const (
	TYPE_UUID           = `Uuid`
	TYPE_AUTO_INCREMENT = `AutoIncrement`
	TYPE_PLAIN          = `Plain`
)

var generatePKStub = `// GeneratePK generate the primary key in {{table}} table
func (m *{{model}}) GeneratePK() {
	m.GenerateId(m)
}`

func generateStructureReplacements(structure *Structure, domain string) ([]string, []string) {
	// Handle timestamps
	timestamps := "false"
	if structure.Timestamps {
		timestamps = "true"
	}

	// Handle softDeletes
	softDeletes := "false"
	if structure.SoftDelete {
		softDeletes = "true"
	}

	// Handle fields and colFields
	fields := ""
	colFields := ""
	jsonFields := ""
	modelFields := ""
	for _, field := range structure.Fields {
		pascalName := util.SnakeToPascalCase(field.Name)
		camelName := util.SnakeToCamelCase(field.Name)

		fields = fmt.Sprintf("%s %s %s `json:\"%s\" db:\"%s\"`\n", fields, pascalName, field.Type, camelName, field.Name)
		jsonFields = fmt.Sprintf("%s %s %s `json:\"%s\"`\n", jsonFields, pascalName, field.Type, camelName)
		colFields = fmt.Sprintf("%s fazzdb.Col(\"%s\"),\n", colFields, field.Name)
		modelFields = fmt.Sprintf("%s m.%s = payload.%s\n", modelFields, pascalName, pascalName)
	}

	// Handle generatePK
	generatePK := ""
	if structure.Type == TYPE_UUID {
		generatePK = generatePKStub
	}

	// Handle names
	camelModel := util.PascalToCamelCase(structure.Model)

	return []string{
			"{{generatePK}}",
			"{{fields}}",
			"{{colFields}}",
			"{{jsonFields}}",
			"{{modelFields}}",
			"{{model}}",
			"{{table}}",
			"{{idType}}",
			"{{idDataType}}",
			"{{timestamps}}",
			"{{softDeletes}}",
			"{{camelModel}}",
			"{{domain}}",
		}, []string{
			generatePK,
			fields,
			colFields,
			jsonFields,
			modelFields,
			structure.Model,
			structure.Table,
			structure.Type,
			structure.DataType,
			timestamps,
			softDeletes,
			camelModel,
			domain,
		}
}
