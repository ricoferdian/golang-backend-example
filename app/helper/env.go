package helper

import "os"

var (
	developmentEnv = "development"
	//stagingEnv     = "staging"
	//productionEnv  = "production"
)

func Getenv() string {
	env := os.Getenv("APP_ENVIRONMENT")
	if env == "" {
		return developmentEnv
	}
	return env
}

func GetConfigPath(env string) string {
	if env == developmentEnv {
		return "files/etc/config"
	}
	return "files/etc/config"
}

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown-hostname"
	}
	return hostname
}
