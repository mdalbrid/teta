package main

type config struct {
	User     string `required:"true" default:"postgres"`
	Password string `required:"false" default:""`
	DBName   string `required:"true" default:"postgres"`
	Host     string `required:"true" default:"localhost"`
	Port     string `required:"true" default:"5432"`
}
