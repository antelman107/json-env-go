package decode

import (
	"encoding/json"
	"os"
)

func DecodeConfigFromEnv(envName string, target interface{}) error {
	return json.Unmarshal([]byte(os.Getenv(envName)), &target)
}
