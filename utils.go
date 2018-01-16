package main

import "os"

// GetEnv returns environment value for given key
func GetEnv(key string, def string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	return def
}
