package helper

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type AppConfig struct {
	ServerConf     ServerConfig     `yaml:"http_server"`
	DBConf         DatabaseConfig   `yaml:"database"`
	RediConf       RedisConfig      `yaml:"redis"`
	MonitoringConf MonitoringConfig `yaml:"monitoring"`
	JWTConf        JWTConfig        `yaml:"jwt_auth"`
}

type JWTConfig struct {
	HeaderKey      string `yaml:"header_key"`
	ClaimIssuer    string `yaml:"claim_issuer"`
	ClaimAudience  string `yaml:"claim_audience"`
	ClaimExpirySec int    `yaml:"claim_expiry_sec"` // in seconds
}

type MonitoringConfig struct {
	NewRelicKey              string `yaml:"newrelic_key"`
	EnableLogForwarding      bool   `yaml:"enable_log_forwarding"`
	EnableDistributedTracing bool   `yaml:"enable_distributed_tracing"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	DriverName string `yaml:"driver_name"`
	Hostname   string `yaml:"hostname"`
	Port       string `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"db_name"`
}

type RedisConfig struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
}

func InitConfig(appName string) (config *AppConfig) {
	env := Getenv()
	configPath := GetConfigPath(env)
	fPath := fmt.Sprintf("%s/%s.%s.yaml", configPath, appName, env)
	err := ReadConfig(&config, fPath)
	if err != nil {
		log.Printf("Read config returned err : %s\n", err.Error())
		log.Fatalln(fmt.Sprintf("Fatal config file not found in path %s", fPath))
		return
	}
	return config
}

func ReadConfig(dest interface{}, path string) error {
	// check if this path is exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, dest)
}
