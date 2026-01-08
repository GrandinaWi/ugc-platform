package config

import "os"

type Config struct {
	Port      string
	JWTSecret []byte
	UsersAPI  string
}

func Load() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		JWTSecret: []byte(getEnv("JWT_SECRET", "ugc-secret")),
		UsersAPI:  getEnv("USERS_API", "http://users-api:8080"),
	}
}
func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
