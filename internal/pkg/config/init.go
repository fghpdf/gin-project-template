package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Init use viper to read the root path /configs xxx.toml file.
// First use this function to init viper
// Then use viper.Get('xxxx') can get the variable
func Init() {
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("default")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
