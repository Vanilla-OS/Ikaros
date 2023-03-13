package core

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PkgManager string `json:"pkgmanager"`
	Debug      bool   `json:"debug"`
}

var Cnf *Config

func init() {
	viper.AddConfigPath("/usr/share/ikaros/")
	viper.AddConfigPath("config/")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&Cnf)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func GetPkgManager() string {
	return Cnf.PkgManager
}

func GetDebug() bool {
	return Cnf.Debug
}
