package decode

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	envName := "CONFIG"
	env := `{"a":"1","b":"2"}`

	type config struct {
		A string `json:"a"`
		B string `json:"b"`
	}

	os.Setenv(envName, env)

	var a config
	if err := DecodeConfigFromEnv(envName, &a); err != nil {
		panic(err)
	}

	assert.Equal(t, config{ A: "1", B:"2"}, a)
}