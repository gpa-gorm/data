package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type EntryArgsType struct {
	Run       bool `arg:"-r,help:Run the server"`
	Seed      bool `arg:"--seed,help:Seed the database"`
	SyncForce bool `arg:"--syncf,help:Sync Force the database"`
}

type AllConfig struct {
	DB    dbConfig
	App   appConfig
	Other otherConfig
}

var (
	DBAppCfg        dbConfig
	AppCfg          appConfig
	EntryArgsAppCfg EntryArgsType
	OtherAppCfg     otherConfig
)

// Validate validates the configuration values
func validateMultipleStruct(structs ...interface{}) {
	validate := validator.New()

	for _, s := range structs {
		if err := validate.Struct(s); err != nil {
			panic(err)
		}
	}
}

// Validates and loads the configuration values
func Validate() error {

	var configuration *AllConfig

	// get the root path of the project
	_, file, _, _ := runtime.Caller(0)
	//rootPath := path.Join(file, "..", "..", "..")
	rootPath := path.Join(file, "..", "..")

	fmt.Println("Project root path: ", rootPath)

	// set the config file
	envPath := path.Join(rootPath, ".env")
	// viper.AddConfigPath
	// viper.SetConfigName
	// viper.SetConfigType
	viper.SetConfigFile(envPath)

	if err := viper.ReadInConfig(); err != nil {
		//log.Errorf("Error to reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		//log.Errorf("error to decode, %v", err)
		return err
	}

	DBAppCfg = dbConfig{
		Host:     viper.Get("DB_HOST").(string),
		Port:     viper.Get("DB_PORT").(string),
		User:     viper.Get("DB_USER").(string),
		Password: viper.Get("DB_PASS").(string),
		DBName:   viper.Get("DB_NAME").(string),
	}

	AppCfg = appConfig{
		Port:        viper.Get("PORT").(string),
		Environment: viper.Get("ENV").(string),
		//SSL:          viper.Get("SSL") == "TRUE",
		//IsProduction: viper.Get("ENV") == "PRODUCTION",
		APIVersion: viper.Get("API_VERSION").(string),
	}

	validateMultipleStruct(DBAppCfg, AppCfg, OtherAppCfg)

	return nil
}
