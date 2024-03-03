package validation

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func EnvCheck(config interface{}, fileENV ...string) {

	godotenv.Load(fileENV...)

	configType := reflect.TypeOf(config).Elem()
	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		envKey := field.Tag.Get("env")

		if envKey == "" {
			log.Fatalf("Config: env tag not found for field %s", field.Name)
		}

		envValue, found := os.LookupEnv(envKey)
		if !found {
			log.Fatalf("Config: %s environment variable not found", envKey)
		}

		fieldValue := configValue.Field(i)
		kind := fieldValue.Kind()

		switch kind {
		case reflect.String:
			fieldValue.SetString(envValue)
		case reflect.Int:
			v, err := strconv.Atoi(envValue)
			if err != nil {
				logrus.Fatalf("Config: invalid %s value, %s", envKey, err.Error())
			}
			fieldValue.SetInt(int64(v))
		default:
			logrus.Fatalf("Config: unsupported field type %s for %s", fieldValue.Kind(), envKey)
		}
	}
}
