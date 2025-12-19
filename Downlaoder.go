package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type ImageDownloader struct {
	config *Config
	client *http.Client
}

func NewImageDownloader(cfg *Config) *ImageDownloader {
	return &ImageDownloader{
		config: cfg,
		client: &http.Client{
			Timeout: cfg.RequestTimeout,
		},
	}
}

func (d *ImageDownloader) GetGoogleImageURLs(query string, count int) ([]string, error) {
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s&tbm=isch", url.QueryEscape(query))

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return d.extractImageURLs(string(body), count), nil
}

func (d *ImageDownloader) extractImageURLs(html string, count int) []string {
	var urls []string

	patterns := []string{
		`"ou":"(https?://[^"]+)"`,
		`\["(https?://[^"]+\.(?:jpg|jpeg|png|gif|webp)[^"]*)"`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(html, -1)

		for _, match := range matches {
			if len(match) > 1 {
				imgURL := match[1]
				if !strings.Contains(imgURL, "gstatic.com") &&
					!strings.Contains(imgURL, "google.com") &&
					len(imgURL) > 20 {
					urls = append(urls, imgURL)
					if len(urls) >= count {
						return urls
					}
				}
			}
		}
	}

	return urls
}

func (d *ImageDownloader) DownloadImage(imgURL, filename string) error {
	req, err := http.NewRequest("GET", imgURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := d.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status: %d", resp.StatusCode)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
