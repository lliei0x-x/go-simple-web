package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

func getPostgreConfg() string {
	var (
		dbName   string
		port     string
		user     string
		sslMode  string
		password string
		host     string
	)

	dbName = viper.GetString("db.dbname")
	port = viper.GetString("db.port")
	user = viper.GetString("db.user")
	sslMode = viper.GetString("db.sslmode")
	password = viper.GetString("db.password")
	host = viper.GetString("db.host")

	return fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s", host, user, dbName, port, sslMode, password)
}

// GetSessionConfig func
func GetSessionConfig() (store string, sessionName string) {
	store = viper.GetString("session.store")
	sessionName = viper.GetString("session.sessionName")
	return
}
