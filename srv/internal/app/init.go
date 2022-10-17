package app

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

func NewConfig(user, password, DBName, Host, Port string) *config {
	return &config{
		user,
		password,
		DBName,
		Host,
		Port,
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (c *config) NewPostgresDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable ",
			c.User, c.Password, c.DBName, c.Host, c.Port))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
