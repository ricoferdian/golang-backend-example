package helper

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type AppConfig struct {
	ServerConf           *ServerConfig         `yaml:"http_server"`
	HandlerConf          *HandlerConfig        `yaml:"http_handler"`
	DBConf               *DatabaseConfig       `yaml:"database"`
	RediConf             *RedisConfig          `yaml:"redis"`
	MonitoringConf       *MonitoringConfig     `yaml:"monitoring"`
	JWTConf              *JWTConfig            `yaml:"jwt_auth"`
	StoreKitVerifyConfig *StoreKitVerifyConfig `yaml:"storekit_verify"`
	SecureOtpConfig      *SecureOtpConfig      `yaml:"secure_otp"`
	AWSConfig            *AWSConfig            `yaml:"aws"`
}

type JWTConfig struct {
	HeaderKey      string `yaml:"header_key"`
	ClaimIssuer    string `yaml:"claim_issuer"`
	ClaimAudience  string `yaml:"claim_audience"`
	ClaimExpirySec int    `yaml:"claim_expiry_sec"` // in seconds
}

type SecureOtpConfig struct {
	OtpIssuer    string `yaml:"otp_issuer"`
	OtpExpirySec int    `yaml:"otp_expiry_sec"` // in seconds
}

type MonitoringConfig struct {
	NewRelicKey              string `yaml:"newrelic_key"`
	EnableLogForwarding      bool   `yaml:"enable_log_forwarding"`
	EnableDistributedTracing bool   `yaml:"enable_distributed_tracing"`
	EnableCodeLevelMetrics   bool   `yaml:"enable_code_level_metrics"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type HandlerConfig struct {
	Timeout int `yaml:"timeout"`
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

type StoreKitVerifyConfig struct {
	Hostname string `yaml:"hostname"`
	Timeout  int    `yaml:"timeout"`
}

type AWSConfig struct {
	S3Config *S3BucketConfig   `yaml:"s3"`
	CFConfig *CloudFrontConfig `yaml:"cloudfront"`
}

type CloudFrontConfig struct {
	URL string `yaml:"url"`
}

type S3BucketConfig struct {
	BucketName    string               `yaml:"bucket_name"`
	ContentConfig *KoraContentS3Config `yaml:"content"`
}

type KoraContentS3Config struct {
	Path string `yaml:"path"`
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
