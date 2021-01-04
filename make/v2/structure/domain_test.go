package structure_test

import (
	"github.com/payfazz/tango/make/v2/structure"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

func TestDomainUnmarshall(t *testing.T) {
	file, err := os.Open("../example/domain-note.yaml")
	require.NoError(t, err)

	byteVal, err := ioutil.ReadAll(file)
	require.NoError(t, err)

	var domain structure.Domain
	err = yaml.Unmarshal(byteVal, &domain)
	require.NoError(t, err)

	require.NotEmpty(t, domain.Package)
	require.NotNil(t, domain.Service)
	require.NotNil(t, domain.Package)
}
