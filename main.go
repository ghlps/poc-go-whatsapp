package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"
)

func main() {
	cfg := Load()

	if cfg.IsDev {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		data, err := os.ReadFile("event.json")
		if err != nil {
			log.Fatalf("failed to read event.json: %v", err)
		}

		_, err = handler(ctx, data)
		if err != nil {
			fmt.Printf("handler error: %v\n", err)
			os.Exit(1)
		}
	} else {
		lambda.Start(handler)
	}
}
