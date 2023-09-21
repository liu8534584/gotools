package gotools

import "os"

const (
	EnvDevelop    = "dev"
	EnvTest       = "test"
	EnvProduction = "prod"
)

func GetApplicationEnvInfo() string {
	var envKey = "APPLICATION_ENV"
	var envVal = os.Getenv(envKey)
	if envVal == "" {
		return EnvDevelop
	}
	return envVal
}
