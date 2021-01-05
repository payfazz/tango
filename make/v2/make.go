package make

import (
	"fmt"
	"github.com/payfazz/tango/make/v2/structure"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type Generator func(path string, content []byte) error

func Generate(path string, content []byte, base structure.Base) error {
	generatorMap := map[string]Generator{
		"group":  generateGroup,
		"domain": generateDomain,
	}
	f, ok := generatorMap[base.Type]
	if !ok {
		return fmt.Errorf("type %s not suppported", base.Type)
	}
	return f(path, content)
}

func generateGroup(path string, content []byte) error {
	dir, _ := filepath.Split(path)

	var group structure.Group
	err := yaml.Unmarshal(content, &group)
	if nil != err {
		return err
	}

	for _, res := range group.Resources {
		filePath := filepath.Join(dir, res)
		resContent, err := ioutil.ReadFile(filePath)
		if nil != err {
			return err
		}

		var structureBase structure.Base
		err = yaml.Unmarshal(resContent, &structureBase)
		if nil != err {
			return err
		}

		err = Generate(filePath, resContent, structureBase)
		if nil != err {
			return err
		}
	}
	return nil
}

func generateDomain(_ string, content []byte) error {
	var domain structure.Domain
	err := yaml.Unmarshal(content, &domain)
	if nil != err {
		return err
	}
	err = domain.GenerateDefault()
	if nil != err {
		return err
	}

	homeDir, _ := os.UserHomeDir()
	tangoDir := filepath.Join(homeDir, "/.tango")
	tpl := template.Must(template.ParseGlob(filepath.Join(tangoDir, BASE_STUB_FILE_GLOB)))
	tpl = template.Must(tpl.ParseGlob(filepath.Join(tangoDir, METHOD_STUB_FILE_GLOB)))

	return generateDomainFile(filepath.Join(DOMAIN_DIR, domain.Package), domain, tpl)
}

func generateDomainFile(domainDir string, domain structure.Domain, tpl *template.Template) error {
	for _, subdomain := range domain.SubDomains {
		err := generateDomainFile(filepath.Join(domainDir, subdomain.Package), subdomain, tpl)
		if err != nil {
			return err
		}
	}
	if domain.Model != nil {
		modelFile, err := createFileIfNotExist(filepath.Join(domainDir, "model.go"))
		if nil != err {
			return err
		}
		defer modelFile.Close()

		err = tpl.ExecuteTemplate(modelFile, "model", domain)
		if nil != err {
			return err
		}
	}
	if domain.Service != nil {
		serviceFile, err := createFileIfNotExist(filepath.Join(domainDir, "service.go"))
		if nil != err {
			return err
		}
		defer serviceFile.Close()
		err = tpl.ExecuteTemplate(serviceFile, "service", domain)
		if nil != err {
			return err
		}

		payloadFile, err := createFileIfNotExist(filepath.Join(domainDir, "payload.go"))
		if nil != err {
			return err
		}
		defer payloadFile.Close()
		err = tpl.ExecuteTemplate(payloadFile, "payload", domain)
		if nil != err {
			return err
		}

		resultFile, err := createFileIfNotExist(filepath.Join(domainDir, "result.go"))
		if nil != err {
			return err
		}
		defer resultFile.Close()
		err = tpl.ExecuteTemplate(resultFile, "result", domain)
		if nil != err {
			return err
		}

		// generate methods
		for _, meth := range domain.Service.MethodImpls {
			err = tpl.ExecuteTemplate(serviceFile, fmt.Sprintf("method_%s", meth.Type), meth)
			if nil != err {
				return err
			}
			err = tpl.ExecuteTemplate(payloadFile, fmt.Sprintf("payload_%s", meth.Type), meth)
			if nil != err {
				return err
			}
			err = tpl.ExecuteTemplate(resultFile, fmt.Sprintf("result_%s", meth.Type), meth)
			if nil != err {
				return err
			}
		}
		for _, meth := range domain.Service.SubdomainMethodImpls {
			err = tpl.ExecuteTemplate(serviceFile, fmt.Sprintf("method_%s", meth.Type), meth)
			if nil != err {
				return err
			}
			err = tpl.ExecuteTemplate(payloadFile, fmt.Sprintf("payload_%s", meth.Type), meth)
			if nil != err {
				return err
			}
			err = tpl.ExecuteTemplate(resultFile, fmt.Sprintf("result_%s", meth.Type), meth)
			if nil != err {
				return err
			}
		}
	}
	return nil
}

func createFileIfNotExist(filePath string) (*os.File, error) {
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s already exists", filePath)
	}

	dir, _ := filepath.Split(filePath)

	err = os.MkdirAll(dir, DIR_FILE_MODE)
	if nil != err {
		return nil, err
	}

	return os.Create(filePath)
}
