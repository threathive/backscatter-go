package config

import (
        "fmt"
        "github.com/fsnotify/fsnotify"
        "github.com/spf13/viper"
)

type Constants struct {
        BackScatter struct {
            ApiServer string
            ApiKey string
        }

}

type Config struct {
        Constants
}


// NewConfig is used to generate a configuration instance which will be passed around the codebase
func New() (*Config, error) {
        config := Config{}
        constants, err := initViper()
        config.Constants = constants
        if err != nil {
                return &config, err
        }

        return &config, err
}

func initViper() (Constants, error) {
        viper.SetConfigName("api.config") // Configuration fileName without the .TOML or .YAML extension
        viper.AddConfigPath(".")           // Search the root directory for the configuration file
        err := viper.ReadInConfig()        // Find and read the config file
        if err != nil {                    // Handle errors reading the config file
                return Constants{}, err
        }
        viper.WatchConfig() // Watch for changes to the configuration file and recompile
        viper.OnConfigChange(func(e fsnotify.Event) {
                fmt.Println("Config file changed:", e.Name)
        })

        var constants Constants
        err = viper.Unmarshal(&constants)
        return constants, err
}

