package make_test

import (
	"fmt"
	"github.com/payfazz/tango/make/v2"
	"github.com/payfazz/tango/make/v2/structure"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
	"text/template"
)

func TestMake(t *testing.T) {
	file, err := os.Open("example/domain-project.yaml")
	require.NoError(t, err)

	byteVal, err := ioutil.ReadAll(file)
	require.NoError(t, err)

	var domain structure.Domain
	err = yaml.Unmarshal(byteVal, &domain)
	require.NoError(t, err)
	err = domain.GenerateDefault()
	require.NoError(t, err)

	tpl := template.Must(template.ParseGlob("template/*.stub"))
	tpl = template.Must(tpl.ParseGlob("template/methods/*.stub"))
	err = os.MkdirAll("./test", make.DIR_FILE_MODE)
	require.NoError(t, err)
	modelFile, err := os.Create("./test/model.go")
	require.NoError(t, err)
	defer modelFile.Close()

	serviceFile, err := os.Create("./test/service.go")
	require.NoError(t, err)
	defer serviceFile.Close()

	payloadFile, err := os.Create("./test/payload.go")
	require.NoError(t, err)
	defer payloadFile.Close()

	resultFile, err := os.Create("./test/result.go")
	require.NoError(t, err)
	defer resultFile.Close()

	err = tpl.ExecuteTemplate(modelFile, "model", domain)
	require.NoError(t, err)
	err = tpl.ExecuteTemplate(serviceFile, "service", domain)
	require.NoError(t, err)
	err = tpl.ExecuteTemplate(payloadFile, "payload", domain)
	require.NoError(t, err)
	err = tpl.ExecuteTemplate(resultFile, "result", domain)
	require.NoError(t, err)

	for _, meth := range domain.Service.MethodImpls {
		err = tpl.ExecuteTemplate(serviceFile, fmt.Sprintf("method_%s", meth.Type), meth)
		require.NoError(t, err)
		err = tpl.ExecuteTemplate(payloadFile, fmt.Sprintf("payload_%s", meth.Type), meth)
		require.NoError(t, err)
		err = tpl.ExecuteTemplate(resultFile, fmt.Sprintf("result_%s", meth.Type), meth)
		require.NoError(t, err)
	}
	for _, meth := range domain.Service.SubdomainMethodImpls {
		err = tpl.ExecuteTemplate(serviceFile, fmt.Sprintf("method_%s", meth.Type), meth)
		require.NoError(t, err)
		err = tpl.ExecuteTemplate(payloadFile, fmt.Sprintf("payload_%s", meth.Type), meth)
		require.NoError(t, err)
		err = tpl.ExecuteTemplate(resultFile, fmt.Sprintf("result_%s", meth.Type), meth)
		require.NoError(t, err)
	}
}
