package config

import (
	"fmt"
)

type dbConfig struct {
	// Driver   string `validate:"required"`
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DBName   string `validate:"required"`
	// LogMode  bool   `validate:"required"`
}

func DbConfiguration(dbConfig *dbConfig) (string, string) {

	//replicaDBName := viper.GetString("REPLICA_DB_NAME")
	//replicaDBUser := viper.GetString("REPLICA_DB_USER")
	//replicaDBPassword := viper.GetString("REPLICA_DB_PASSWORD")
	//replicaDBHost := viper.GetString("REPLICA_DB_HOST")
	//replicaDBPort := viper.GetString("REPLICA_DB_PORT")
	//replicaDBSslMode := viper.GetString("REPLICA_SSL_MODE")

	masterDBDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port, "disable",
	)

	//replicaDBDSN := fmt.Sprintf(
	//	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	//	replicaDBHost, replicaDBUser, replicaDBPassword, replicaDBName, replicaDBPort, replicaDBSslMode,
	//)
	// return masterDBDSN, replicaDBDSN
	return masterDBDSN, "dummy"
}
