package configs

import (
	"strings"

	"github.com/spf13/viper"
)

var Properties ConfigStruct

// InitConfig : initial config
func InitConfig(configPath, serviceName string) error {
	viper.AddConfigPath("./")
	viper.SetConfigName(".configs")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	envPrefix := "RINTIK_AUTH"

	if serviceName != "" {
		envPrefix = strings.ToUpper(serviceName)
		envPrefix = strings.ReplaceAll(envPrefix, " ", "_")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.Unmarshal(&Properties)
	if err != nil {
		return err
	}

	return nil
}
