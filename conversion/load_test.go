package conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfigValue(t *testing.T) {
	values, err := LoadConfigValue("../fixtures/some.tfvars")

	assert.Nil(t, err)
	assert.Equal(t, len(values), 4, "we should have 4 return values")
}

func TestLoadConfigValues(t *testing.T) {
	values, _ := LoadConfigValue("../fixtures/some.tfvars")

	assert.Equal(t, "some_value", values[0].Name)
	assert.Equal(t, "a value", values[0].Value)

	assert.Equal(t, "some_ip", values[1].Name)
	assert.Equal(t, "172.20.50.0/25", values[1].Value)

	assert.Equal(t, "some_list", values[2].Name)
	assert.Equal(t, "172.20.50.0/24,172.20.51.0/24", values[2].Value)

	assert.Equal(t, "some_num", values[3].Name)
	assert.Equal(t, "0", values[3].Value)
}
