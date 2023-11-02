package env

import (
	"github.com/joho/godotenv"
	"strconv"
	"strings"
)

var env map[string]string

func get(key string) string {
	if env == nil {
		e, err := godotenv.Read()
		if err != nil {
			panic("error loading env: " + err.Error())
		}

		env = e
	}

	return env[key]
}

func GetString(key string) string {
	return get(key)
}

func GetInt(key string) int {
	value, err := strconv.Atoi(get(key))
	if err != nil {
		return 0
	}

	return value
}

func GetSlice(key string) []string {
	value := get(key)
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
