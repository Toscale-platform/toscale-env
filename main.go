package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env:", err)
	}
}

func GetString(key string) string {
	return os.Getenv(key)
}

func GetSlice(key string) []string {
	return strings.Split(os.Getenv(key), ",")
}

func GetInt(key string) int {
	num, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 0
	}

	return num
}
