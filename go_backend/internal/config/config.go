package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
    ServerPort     string
    DBHost         string
    DBPort         string
    DBUser         string
    DBPassword     string
    DBName         string
    RedisAddr      string
    RedisPassword  string
    JWTSecret      string
}

func Load() *Config {
	// Try to load .env from multiple possible locations
	
	// Get current working directory first
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		log.Printf("Error getting working directory: %v", wdErr)
	}
	
	// Build list of paths to try
	pathsToTry := []string{
		".env", // Current directory (most common case)
	}
	
	// Add relative paths
	pathsToTry = append(pathsToTry,
		filepath.Join("..", ".env"),               // One level up
		filepath.Join("..", "..", ".env"),         // Two levels up (for cdm/api)
		filepath.Join("..", "..", "..", ".env"),   // Three levels up
	)
	
	// Add absolute paths if we have working directory
	if wdErr == nil {
		// Current directory absolute path
		pathsToTry = append(pathsToTry, filepath.Join(wd, ".env"))
		
		// If in cdm/api, go up to go_backend
		if filepath.Base(wd) == "api" {
			parent := filepath.Dir(wd)
			if filepath.Base(parent) == "cdm" {
				goBackendRoot := filepath.Dir(parent)
				pathsToTry = append(pathsToTry, filepath.Join(goBackendRoot, ".env"))
			}
		}
	}
	
	// Try each path
	found := false
	for _, path := range pathsToTry {
		// Check if file exists first
		if _, statErr := os.Stat(path); statErr == nil {
			// Read file and remove BOM if present
			fileContent, readErr := ioutil.ReadFile(path)
			if readErr != nil {
				log.Printf("Failed to read .env from %s: %v", path, readErr)
				continue
			}
			
			// Remove UTF-8 BOM if present
			fileContent = bytes.TrimPrefix(fileContent, []byte("\xef\xbb\xbf"))
			
			// Parse the content
			envMap, parseErr := godotenv.Parse(bytes.NewReader(fileContent))
			if parseErr != nil {
				log.Printf("Failed to parse .env from %s: %v", path, parseErr)
				continue
			}
			
			// Set environment variables
			for key, value := range envMap {
				os.Setenv(key, value)
			}
			
			found = true
			log.Printf("Loaded .env from %s", path)
			break
		}
	}
	
	if !found {
		log.Printf("No .env file found or could not be loaded. Tried: %v. Using system environment variables.", pathsToTry)
	}
    return &Config{
        ServerPort:    os.Getenv("SERVER_PORT"),
        DBHost:        os.Getenv("DB_HOST"),
        DBPort:        os.Getenv("DB_PORT"),
        DBUser:        os.Getenv("DB_USER"),
        DBPassword:    os.Getenv("DB_PASSWORD"),
        DBName:        os.Getenv("DB_NAME"),
        RedisAddr:     os.Getenv("REDIS_ADDR"),
        RedisPassword: os.Getenv("REDIS_PASSWORD"),
        JWTSecret:     os.Getenv("JWT_SECRET"),
    }
}