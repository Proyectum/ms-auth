package boot

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"regexp"
	"time"
)

var CONFIG Configuration

type ApplicationConf struct {
	Name string `yaml:"name"`
	Env  string `yaml:"environment"`
}

type JWTConf struct {
	Secret     string        `yaml:"secret"`
	Expiration time.Duration `yaml:"expiration"`
}

type PasswordConfig struct {
	Cost int `yaml:"cost"`
}

type SecurityConf struct {
	JWT      JWTConf        `yaml:"jwt"`
	Password PasswordConfig `yaml:"password"`
}

type PostgresConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type DatasourceConf struct {
	Postgres PostgresConf `yaml:"postgres"`
}

type DataConf struct {
	Datasource DatasourceConf `yaml:"datasource"`
}

type ServerConf struct {
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
}

type Configuration struct {
	Application ApplicationConf `yaml:"app"`
	Security    SecurityConf    `yaml:"security"`
	Data        DataConf        `yaml:"data"`
	Server      ServerConf      `yaml:"server"`
}

func LoadConfig() {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "standalone"
	}
	viper.SetConfigName(fmt.Sprintf("application-%s.yaml", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	replaceEnvVariables()

	if err := viper.Unmarshal(&CONFIG, func(decoderConf *mapstructure.DecoderConfig) {
		decoderConf.TagName = "yaml"
		decoderConf.ErrorUnset = true
	}); err != nil {
		panic(fmt.Errorf("falta error unmarshal config: %w", err))
	}

}

func replaceEnvVariables() {
	var replacePlaceholders func(map[string]interface{}) map[string]interface{}

	replacePlaceholders = func(m map[string]interface{}) map[string]interface{} {
		for k, v := range m {
			switch v := v.(type) {
			case map[string]interface{}:
				m[k] = replacePlaceholders(v)
			case string:
				m[k] = replaceString(v)
			}
		}
		return m
	}

	settings := replacePlaceholders(viper.AllSettings())

	if err := viper.MergeConfigMap(settings); err != nil {
		panic(fmt.Errorf("falta error mergin config: %w", err))
	}
}

func replaceString(s string) string {
	re := regexp.MustCompile(`\$\{(\w+)(?::([^}]*))?}`)
	return re.ReplaceAllStringFunc(s, func(matched string) string {
		parts := re.FindStringSubmatch(matched)
		envVar := parts[1]
		defaultValue := parts[2]

		if value, exists := os.LookupEnv(envVar); exists {
			return value
		}
		return defaultValue
	})
}
