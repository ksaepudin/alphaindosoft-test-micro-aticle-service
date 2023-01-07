package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ksaepudin/alphaindosoft-test-micro-aticle-service/config/db"
	"github.com/spf13/viper"
)

type config struct {
	Database db.DatabaseList
}

var cfg config
var (
	_, b, _, _ = runtime.Caller(0)
	dir        = filepath.Dir(b)
)

func init() {
	_, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(dir + "/db")
	viper.SetConfigType("yaml")
	viper.SetConfigName("db.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load database config: %v", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	viper.Unmarshal(&cfg)
	dataByte, _ := json.Marshal(&cfg)

	fmt.Println("=============================")
	fmt.Println(string(dataByte))
	fmt.Println("=============================")
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(env) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}

func GetConfig() *config {
	return &cfg
}
