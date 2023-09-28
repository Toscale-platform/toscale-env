package env

import (
	"github.com/joho/godotenv"
	"strconv"
	"strings"
)

var env = make(map[string]string)

func Init() {
	var err error

	env, err = godotenv.Read()
	if err != nil {
		panic("error loading env: " + err.Error())
	}
}

func GetString(key string) string {
	return env[key]
}

func GetInt(key string) int {
	value, err := strconv.Atoi(env[key])
	if err != nil {
		return 0
	}

	return value
}

func GetSlice(key string) []string {
	value := env[key]

	if value == "" {
		return []string{}
	}

	return strings.Split(value, ",")
}

func GetBool(key string) bool {
	value, err := strconv.ParseBool(env[key])
	if err != nil {
		return false
	}

	return value
}
