package helpers

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
	// 	err := godotenv.Load(".env.production")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// } else if os.Getenv("ENVIRONMENT") == "STAGING" {
	// 	err := godotenv.Load(".env.staging")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// } else {
	// 	err := godotenv.Load(".env.local")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	return os.Getenv(key)
}
func ContextTimeOut() (timeOut int) {
	val := LoadEnv("CONTEXT_TIME_OUT")
	num, err := strconv.Atoi(val)
	if err != nil {
		return 10
	}
	return num

}
