package main

import (
	"os"
	"time"
)

type Config struct {
	// Database configuration
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	// Image download configuration
	NumImages      int           // Number of images to download per product
	ChunkSize      int           // Number of products to process per chunk
	ChunkDelay     time.Duration // Delay between chunks
	ProductDelay   time.Duration // Delay between products
	FolderName     string        // Folder to save images
	RequestTimeout time.Duration // HTTP request timeout
}

func NewConfig() *Config {
	return &Config{
		DBUser:         getEnv("DB_USER", "appuser"),
		DBPassword:     getEnv("DB_PASSWORD", "apppassword"),
		DBHost:         getEnv("DB_HOST", "mysql"),
		DBPort:         getEnv("DB_PORT", "3307"),
		DBName:         getEnv("DB_NAME", "products_db"),
		FolderName:     getEnv("FOLDER_NAME", "images"),
		NumImages:      4,
		ChunkSize:      5,
		ChunkDelay:     10 * time.Second,
		ProductDelay:   2 * time.Second,
		RequestTimeout: 30 * time.Second,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
