package main

import (
	"context"
	"fmt"

	"github.com/jmarren/toucan/internal"
	"github.com/jmarren/toucan/internal/models"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	models.Init(context.Background())
	fmt.Println("hi from main")

	internal.Start()
}
