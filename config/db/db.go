package db

type DatabaseList struct {
	Articles struct {
		Mysql Database
	}
}
type Database struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Dbname       string `yaml:"dbname"`
	Adapter      string `yaml:"adapter"`
	MaxOpen      int    `yaml:"maxOpen"`
	MaxIdle      int    `yaml:"maxIdle"`
	IdleLifeTime int    `yaml:"idleLifeTime"`
	OpenLifeTime int    `yaml:"openLifeTime"`
	CustomPool   bool   `yaml:"customPool"`
}
