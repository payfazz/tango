package make

// StructureMap handle overall structure mapping
type StructureMap struct {
	Structures []*Structure `yaml:"structures"`
}

// Structure handle structure.yaml mapping
type Structure struct {
	Model      string       `yaml:"model"`
	Table      string       `yaml:"table"`
	Type       string       `yaml:"idType"`
	DataType   string       `yaml:"idDataType"`
	Timestamps bool         `yaml:"timestamps"`
	SoftDelete bool         `yaml:"softDeletes"`
	Fields     []*Field     `yaml:"fields"`
	Action     *Action      `yaml:"action"`
	Childs     []*Structure `yaml:"childs"`
}

// Field handle field inside structure.yaml mapping
type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
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
