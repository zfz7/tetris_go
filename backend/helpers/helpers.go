package helpers

import "os"

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != ""{
		return value
	}
	return fallback
}