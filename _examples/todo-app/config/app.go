package config

type appConfig struct {
	Environment string `validate:"required,oneof=DEV PRODUCTION LOCAL"`
	// Host         string `validate:"required"`
	Port string `validate:"required"`
	// SSL          bool
	// IsProduction bool
	APIVersion string `validate:"required"`
	// Cors []string `validate:"required"`
	//Cookie struct {
	//	Name string `validate:"required"`
	//	SameSite string `validate:"required"`
	//}

}

//type Config struct {
//	Version         string
//	ServerStartTime time.Time
//	Config          Config
//	Logger          *log.Logger
//	Validator       *validator.Validate
//}
