package structure

import (
	"fmt"
	"github.com/payfazz/tango/util"
	"strings"
	"unicode"
)

type Domain struct {
	Package    string   `yaml:"package"`
	Model      *Model   `yaml:"model,omitempty"`
	Service    *Service `yaml:"service,omitempty"`
	SubDomains []Domain `yaml:"subDomains,omitempty"`
}

func (d *Domain) GenerateDefault() error {
	for _, sub := range d.SubDomains {
		err := sub.GenerateDefault()
		if err != nil {
			return err
		}
		if d.Service != nil {
			for _, methImpl := range sub.Service.MethodImpls {
				var params []MethodParam
				for _, param := range methImpl.Params {
					if unicode.IsUpper([]rune(param.Type)[0]) {
						params = append(params, MethodParam{
							Name: param.Name,
							Type: fmt.Sprintf("%s.%s", sub.Package, param.Type),
						})
					} else {
						params = append(params, param)
					}
				}
				var returns []MethodReturn
				for _, ret := range methImpl.Returns {
					if unicode.IsUpper([]rune(ret.Type)[0]) {
						returns = append(returns, MethodReturn{
							Type: fmt.Sprintf("%s.%s", sub.Package, ret.Type),
						})
					} else if strings.HasPrefix(ret.Type, "*") && unicode.IsUpper([]rune(ret.Type)[1]) {
						returns = append(returns, MethodReturn{
							Type: fmt.Sprintf("*%s.%s", sub.Package, ret.Type[1:]),
						})
					} else if strings.HasPrefix(ret.Type, "[]*") && unicode.IsUpper([]rune(ret.Type)[3]) {
						returns = append(returns, MethodReturn{
							Type: fmt.Sprintf("[]*%s.%s", sub.Package, ret.Type[3:]),
						})
					} else {
						returns = append(returns, ret)
					}
				}

				d.Service.SubdomainMethodImpls = append(d.Service.SubdomainMethodImpls, MethodImpl{
					Name:    methImpl.Name,
					Type:    "subdomain_forward",
					Params:  params,
					Returns: returns,
					Data: map[string]string{
						"package": sub.PascalPackage(),
					},
				})
			}
		}
	}
	if d.Package == "" {
		return fmt.Errorf("package must not be empty")
	}
	if d.Model != nil {
		err := d.generateModelDefault()
		if err != nil {
			return err
		}
	}
	if d.Service != nil {
		err := d.generateServiceDefault()
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Domain) generateModelDefault() error {
	m := d.Model
	if m.Table == "" {
		return fmt.Errorf("package must not be empty")
	}
	if m.Type == "" {
		m.Type = "UuidModel"
	}
	if m.IdName == "" {
		m.IdName = "id"
	}
	if m.IdType == "" {
		if m.Type == "UuidModel" {
			m.IdType = "string"
		} else {
			m.IdType = "int64"
		}
	}
	if m.Timestamps == nil {
		timestampDefault := true
		m.Timestamps = &timestampDefault
	}
	if m.SoftDelete == nil {
		softDeleteDefault := true
		m.SoftDelete = &softDeleteDefault
	}
	return nil
}

func (d *Domain) generateServiceDefault() error {
	s := d.Service
	var methodImpls []MethodImpl
	for _, def := range s.MethodDefs {
		methodImpl := MethodImpl{
			Name: def["name"].(string),
			Type: def["type"].(string),
			Params: []MethodParam{
				{Name: "ctx", Type: "context.Context"},
			},
			Returns: []MethodReturn{
				{Type: "error"},
			},
		}
		switch methodImpl.Type {
		case "model_create":
			methodImpl.Params = append(methodImpl.Params, MethodParam{
				Name: "payload",
				Type: "CreatePayload",
			})
			methodImpl.Payload = Payload{Fields: d.Model.Fields}
		case "model_update":
			methodImpl.Params = append(methodImpl.Params, MethodParam{
				Name: "payload",
				Type: "UpdatePayload",
			})
			methodImpl.Payload = Payload{
				Fields: d.Model.Fields,
			}
			methodImpl.Data = map[string]string{
				"idName":       d.Model.IdName,
				"pascalIdName": d.Model.PascalIdName(),
				"camelIdName":  d.Model.CamelIdName(),
				"idType":       d.Model.IdType,
			}
		case "model_delete":
			methodImpl.Params = append(methodImpl.Params, MethodParam{
				Name: d.Model.IdName,
				Type: d.Model.IdType,
			})
			methodImpl.Data = map[string]string{
				"idName":       d.Model.IdName,
				"pascalIdName": d.Model.PascalIdName(),
				"camelIdName":  d.Model.CamelIdName(),
			}
		case "model_get":
			methodImpl.Params = append(methodImpl.Params, MethodParam{
				Name: d.Model.IdName,
				Type: d.Model.IdType,
			})
			methodImpl.Returns = append([]MethodReturn{{
				Type: "*Model",
			}}, methodImpl.Returns...)
			methodImpl.Data = map[string]string{
				"idName":       d.Model.IdName,
				"pascalIdName": d.Model.PascalIdName(),
				"camelIdName":  d.Model.CamelIdName(),
			}
		case "model_list":
			methodImpl.Params = append(methodImpl.Params,
				MethodParam{
					Name: "payload",
					Type: "ListPayload",
				},
				MethodParam{
					Name: "page",
					Type: "*filter.Page",
				},
			)
			methodImpl.Returns = append([]MethodReturn{{
				Type: "[]*Model",
			}}, methodImpl.Returns...)
		}

		methodImpls = append(methodImpls, methodImpl)
	}
	s.MethodImpls = append(methodImpls, s.MethodImpls...)
	return nil
}

func (d *Domain) PascalPackage() string {
	return util.SnakeToPascalCase(d.Package)
}

type Model struct {
	Table      string  `yaml:"table"`
	Type       string  `yaml:"type"`
	IdName     string  `yaml:"idName"`
	IdType     string  `yaml:"idType"`
	Timestamps *bool   `yaml:"timestamps"`
	SoftDelete *bool   `yaml:"softDelete"`
	Fields     []Field `yaml:"fields"`
}

func (m Model) PascalIdName() string {
	return util.SnakeToPascalCase(m.IdName)
}

func (m Model) CamelIdName() string {
	return util.SnakeToCamelCase(m.IdName)
}

type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func (m Field) PascalName() string {
	return util.SnakeToPascalCase(m.Name)
}

func (m Field) CamelName() string {
	return util.SnakeToCamelCase(m.Name)
}

type Service struct {
	MethodDefs           []MethodDef  `yaml:"methods"`
	MethodImpls          []MethodImpl `yaml:"-"`
	SubdomainMethodImpls []MethodImpl `yaml:"-"`
}

type MethodDef map[string]interface{}

type MethodImpl struct {
	Name    string
	Type    string
	Params  []MethodParam
	Returns []MethodReturn
	Payload Payload
	Result  Result
	Data    map[string]string
}

type Payload struct {
	Fields []Field `yaml:"fields"`
}

type Result struct {
	Fields []Field `yaml:"fields"`
}

type MethodParam struct {
	Name string
	Type string
}

type MethodReturn struct {
	Type string
}
