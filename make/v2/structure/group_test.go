package structure_test

import (
	"github.com/payfazz/tango/make/v2/structure"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

func TestGroupUnmarshall(t *testing.T) {
	file, err := os.Open("../example/main.yaml")
	require.NoError(t, err)

	byteVal, err := ioutil.ReadAll(file)
	require.NoError(t, err)

	var group structure.Group
	err = yaml.Unmarshal(byteVal, &group)
	require.NoError(t, err)
	require.NotEmpty(t, group.Resources)
}
