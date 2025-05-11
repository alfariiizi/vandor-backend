package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func getEnv(key string, validate bool, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if validate && !exists {
		panic("Environment variable " + key + " not set")
	}

	if exists {
		return value
	}

	return defaultValue
}

func getEnvAsInt(key string, validate bool, defaultValue int) int {
	value, exists := os.LookupEnv(key)

	if validate && !exists {
		panic("Environment variable " + key + " not set")
	}

	if exists {
		intValue, err := strconv.Atoi(value)
		if validate && err != nil {
			panic("Environment variable " + key + " is not a valid integer")
		} else if !validate && err != nil {
			log.Println("Environment variable " + key + " is not a valid integer, using default value")
			return defaultValue
		} else if err == nil {
			return intValue
		}
	}

	return defaultValue
}

func getEnvByPrefix(prefix string) map[string]string {
	result := make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		key, value := parts[0], parts[1]
		if strings.HasPrefix(key, prefix) {
			result[key] = value
		}
	}
	return result
}
