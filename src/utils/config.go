package utils

import (
	"flag"
	"log"
)

func InitEnv() {
	// setting environment type
	var environment, envfile string
	flag.StringVar(&environment, "environment", "dev", "Environment (dev, prod, ..)")
	flag.Parse()
	switch environment {
	case "prod":
		Env.Type.Prod = true
		envfile = ".env.prod"
	default:
		Env.Type.Dev = true
		envfile = ".env"
	}

	// setting up .env file
	err := LoadEnvFile(envfile)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", envfile, err)
	}
}
