package main

import (
	"context"

	"github.com/jmarren/toucan/internal"
	"github.com/jmarren/toucan/internal/cache"
	"github.com/jmarren/toucan/internal/models"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	models.Init(context.Background())
	cache.InitCache()
	internal.Start()
}
