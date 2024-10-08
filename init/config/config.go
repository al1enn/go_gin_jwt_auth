package config

import (
    "log"

    "github.com/joho/godotenv"
)

func LoadEnv() {
  err := godotenv.Load("./init/envs/.env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  
}