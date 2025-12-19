package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	// Load configuration
	cfg := NewConfig()

	// Create images folder if it doesn't exist
	if err := os.MkdirAll(cfg.FolderName, 0755); err != nil {
		fmt.Printf("Error creating folder: %v\n", err)
		return
	}

	// Connect to database
	db, err := NewDatabase(cfg)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer db.Close()

	fmt.Println("Connected to database successfully!")

	// Get product names from database
	productNames, err := db.GetProducts()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found %d products to process\n", len(productNames))

	// Initialize image downloader
	downloader := NewImageDownloader(cfg)

	// Process products in chunks
	productCount := 0
	totalChunks := (len(productNames) + cfg.ChunkSize - 1) / cfg.ChunkSize

	for chunkIndex := 0; chunkIndex < len(productNames); chunkIndex += cfg.ChunkSize {
		end := chunkIndex + cfg.ChunkSize
		if end > len(productNames) {
			end = len(productNames)
		}

		currentChunk := (chunkIndex / cfg.ChunkSize) + 1
		fmt.Printf("\n========== Processing Chunk %d/%d ==========\n", currentChunk, totalChunks)

		// Process each product in the chunk
		for i := chunkIndex; i < end; i++ {
			productName := productNames[i]
			productCount++
			fmt.Printf("\n[%d] Processing: %s\n", productCount, productName)

			// Get image URLs from Google Images
			imageURLs, err := downloader.GetGoogleImageURLs(productName.Name, cfg.NumImages)
			if err != nil {
				fmt.Printf("Error getting image URLs: %v\n", err)
				continue
			}

			if len(imageURLs) == 0 {
				fmt.Println("No images found")
				continue
			}

			// Download images
			fmt.Printf("Found %d images, downloading...\n", len(imageURLs))
			for j, imgURL := range imageURLs {
				filename := filepath.Join(cfg.FolderName, fmt.Sprintf("%s_%d.jpg", strconv.Itoa(productName.ID), j+1))
				if err := downloader.DownloadImage(imgURL, filename); err != nil {
					fmt.Printf("Error downloading image %d: %v\n", j+1, err)
					continue
				}
				fmt.Printf("Downloaded: %s\n", filename)
			}

			// Small delay between products to avoid rate limiting
			time.Sleep(cfg.ProductDelay)
		}

		// Delay between chunks (except after the last chunk)
		if end < len(productNames) {
			fmt.Printf("\n⏳ Waiting %v before next chunk...\n", cfg.ChunkDelay)
			time.Sleep(cfg.ChunkDelay)
		}
	}

	fmt.Printf("\n✅ Processing complete! Processed %d products in %d chunks.\n", productCount, totalChunks)
}
