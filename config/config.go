package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	conf *viper.Viper
}

func (c *AppConfig) init(configFile map[string]string) {
	// get build env
	env := os.Getenv("BUILD_ENV")
	if env == "" {
		env = "development"
	}

	if configFile == nil {
		configFile = map[string]string{}
	}

	// default file (required)
	configFile["config"] = "config"

	c.conf = viper.New()
	c.conf.SetDefault("environment", env)

	// Defaults: App
	c.conf.SetDefault("app.name", "default")
	c.conf.SetDefault("app.port", "8000")
	c.conf.SetDefault("app.grpc.port", "8082")
	c.conf.SetDefault("app.grpc.http_route_prefix", "/v1")

	c.conf.SetDefault("app.request.timeout", 100)
	c.conf.SetDefault("app.request.max_conn", 10)
	c.conf.SetDefault("app.request.vol_threshold", 20)
	c.conf.SetDefault("app.request.sleep_window", 5000)
	c.conf.SetDefault("app.request.err_per_threshold", 50)

	// Defaults: DataBase
	c.conf.SetDefault("database.driver", "postgres")
	c.conf.SetDefault("database.host", "localhost")
	c.conf.SetDefault("database.port", "5432")
	c.conf.SetDefault("database.user", "postgres")
	c.conf.SetDefault("database.pass", "postgres")
	c.conf.SetDefault("database.name", "boiler-db")
	c.conf.SetDefault("database.sslmode", "disable")
	c.conf.SetDefault("database.timeout", 60)
	c.conf.SetDefault("database.dsn", "host=%s port=%s dbname=%s user=%s password=%s sslmode=%s")
	c.conf.SetDefault("database.auto_migrate", "off")

	// Conf Env
	c.conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "_", "__")) // APP_DATA__BASE_PASS -> app.data_base.pass
	c.conf.AutomaticEnv()                                              // Automatically load Env variables

	// Conf Files
	c.conf.SetConfigType("yaml") // We're using yaml
	c.conf.SetConfigName(env)    // Search for a config file that matches our environment
	c.conf.AddConfigPath("./")   // look for config in the working directory
	c.conf.ReadInConfig()        // Find and read the config file

	// Read additional files
	for confFile := range configFile {
		c.conf.SetConfigName(confFile)
		c.conf.MergeInConfig()
	}
}

func (c *AppConfig) GetString(key string) string {
	return c.conf.GetString(key)
}

func (c *AppConfig) GetInt(key string) int {
	return c.conf.GetInt(key)
}

func (c *AppConfig) GetBool(key string) bool {
	return c.conf.GetBool(key)
}

func (c *AppConfig) GetStringSlice(key string) []string {
	return c.conf.GetStringSlice(key)
}

func (c *AppConfig) GetStringMap(key string) map[string]interface{} {
	return c.conf.GetStringMap(key)
}

func New(vp *viper.Viper) *AppConfig {
	c := AppConfig{
		conf: vp,
	}
	c.init(map[string]string{"config": "config"})
	return &c
}
