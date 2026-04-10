package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func handler(ctx context.Context) (string, error) {
	bucketName := os.Getenv("S3_BUCKET_NAME")
	dbFileName := os.Getenv("DB_FILE_NAME")
	targetNumber := os.Getenv("TARGET_NUMBER")
	localPath := "/tmp/" + dbFileName

	dbLog := waLog.Stdout("Database", "DEBUG", true)

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config: %w", err)
	}
	s3Client := s3.NewFromConfig(cfg)

	fmt.Println("Downloading database from S3...")
	if err := downloadFromS3(ctx, s3Client, bucketName, dbFileName, localPath); err != nil {
		fmt.Printf("Could not download DB (normal for first run): %v\n", err)
	}

	dbString := fmt.Sprintf("file:%s?_pragma=foreign_keys(1)", localPath)
	container, err := sqlstore.New(ctx, "sqlite", dbString, dbLog)
	if err != nil {
		return "", fmt.Errorf("sqlstore.New: %w", err)
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		return "", fmt.Errorf("GetFirstDevice: %w", err)
	}

	client := whatsmeow.NewClient(deviceStore, waLog.Stdout("Client", "DEBUG", true))

	if client.Store.ID == nil {
		return "", fmt.Errorf("no session found in DB — run QR auth locally first, then upload the DB to S3")
	}

	if err := client.Connect(); err != nil {
		return "", fmt.Errorf("Connect: %w", err)
	}

	if err := sendMessage(client, targetNumber, "Hola desde Go en Lambda!"); err != nil {
		fmt.Printf("Send error: %v\n", err)
	}

	client.Disconnect()

	if err := container.Close(); err != nil {
		fmt.Printf("container.Close error: %v\n", err)
	}

	fmt.Println("Uploading updated database to S3...")
	if err := uploadToS3(ctx, s3Client, bucketName, dbFileName, localPath); err != nil {
		return "", fmt.Errorf("upload to S3: %w", err)
	}

	fmt.Println("Success! Session saved.")
	return "Success", nil
}

func main() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		result, err := handler(ctx)
		if err != nil {
			fmt.Printf("Error en el handler: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Resultado: %s\n", result)
	} else {
		lambda.Start(handler)
	}
}
