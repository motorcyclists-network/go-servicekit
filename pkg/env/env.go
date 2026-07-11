package env

import "os"

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns the provided fallback value.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
