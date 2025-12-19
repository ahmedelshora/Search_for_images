# Go Product Image Downloader

A Go application that automatically downloads product images by connecting to a MySQL database, retrieving product information, and saving images to a local directory.

## Overview

This application:
- Connects to a MySQL database
- Retrieves product IDs and names from `oc_products_descriptions`
- Searches and downloads 4 images per product
- Saves images to a configurable local folder
- Runs containerized using Docker and Docker Compose
- Provides simple management through a Makefile

## Features

- **Go-based**: Fast and efficient image processing
- **Environment Configuration**: Secure credential management via `.env` file
- **Containerized**: Full Docker and Docker Compose support
- **Persistent Storage**: Volume-backed image storage
- **Easy Management**: One-command operations using Makefile

## Project Structure
```
.

└── Main.go
└── config.go
└── Database.go
└── Downloader.go
├── images/                 # Downloaded images storage
├── Dockerfile
├── docker-compose.yml
└── init.sql    # that have the mysql dump you are get to be imported and get the products list 
├── Makefile
├── .env
├── go.mod
└── README.md
```

## Requirements

- Docker
- Docker Compose
- Make

## Setup

### Environment Configuration

Create a `.env` file in the project root:
```env
DB_USER=appuser
DB_PASSWORD=apppassword
DB_HOST=mysql
DB_PORT=3306
DB_NAME=products_db

FOLDER_NAME=images
```

> **Important**: `DB_HOST` must be set to `mysql` when running inside Docker

## Docker Services

The application consists of two services:

1. **app**: Go application container
2. **mysql**: MySQL 8.0 database container

Both images and database data are persisted using Docker volumes.

## Usage

### Makefile Commands

| Command | Description |
|---------|-------------|
| `make up` | Start all services |
| `make down` | Stop all services |
| `make build` | Build and start services |
| `make logs` | Show application logs |
| `make restart` | Restart the Go app |
| `make run` | Run the app once |
| `make clean` | Stop services and remove volumes & images |

### Quick Start

1. **Build and start services**
```bash
   make build
```

2. **Run the application**
```bash
   make run
```

3. **View logs**
```bash
   make logs
```

## Output

- Downloaded images are saved to the `images/` directory
- Each product will have up to 4 downloaded images

## Troubleshooting

### Database Connection Refused

**Solution**: Verify the following in your `.env` file:
```env
DB_HOST=mysql
DB_PORT=3306
```

> **Note**: Inside Docker, services communicate using service names (e.g., `mysql`), not `127.0.0.1` or `localhost`

### Images Not Saving

Check the following:
- Ensure the `images/` directory exists
- Verify Docker volume is mounted correctly in `docker-compose.yml`
- Review application logs: `make logs`
