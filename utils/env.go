package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func setDefaults(env string) {
	if os.Getenv("COPLUK_ENV") == "" {
		os.Setenv("COPLUK_ENV", env)
	}
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}
}

func SetupEnv() {
	env := os.Getenv("COPLUK_ENV")
	if env == "" {
		env = "development"
	}

	if err := godotenv.Load(".env." + env); err != nil {
		panic("Error loading .env." + env + " file")
	}
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	setDefaults(env)
}
