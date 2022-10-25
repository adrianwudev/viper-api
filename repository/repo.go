package repository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	dialect  string
	host     string
	dbport   string
	user     string
	dbname   string
	password string
}

var serverConfig ServerConfig

func GetConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	serverConfig.dialect = viper.GetString("config.server.dialect")
	serverConfig.host = viper.GetString("config.server.host")
	serverConfig.dbport = viper.GetString("config.server.dbport")
	serverConfig.user = viper.GetString("config.server.user")
	serverConfig.dbname = viper.GetString("config.server.dbname")
	serverConfig.password = viper.GetString("config.server.password")
}

func New() {
	GetConfig()

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s  password=%s port=%s", serverConfig.host, serverConfig.user, serverConfig.dbname, serverConfig.password, serverConfig.dbport)
	fmt.Println(dbURI)

	// Openning connection to database
	db, err := gorm.Open(serverConfig.dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("successfully connected to database.")
	}

	// Close connection to database when the main function finishes
	defer db.Close()
}
