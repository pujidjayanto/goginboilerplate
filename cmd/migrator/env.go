package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Ssl      string
}

type Migrator struct {
	Env string
}

type Environment struct {
	Database     Database
	TestDatabase Database
	Migrator     Migrator
}

func loadEnvironment() (*Environment, error) {
	envPath, err := getEnvPath()
	if err != nil || strings.TrimSpace(envPath) == "" {
		return nil, fmt.Errorf("no .env file found")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	return &Environment{
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Ssl:      os.Getenv("DB_SSL_MODE"),
		},
		TestDatabase: Database{
			Host:     os.Getenv("TEST_DB_HOST"),
			Port:     os.Getenv("TEST_DB_PORT"),
			User:     os.Getenv("TEST_DB_USER"),
			Password: os.Getenv("TEST_DB_PASSWORD"),
			Name:     os.Getenv("TEST_DB_NAME"),
			Ssl:      os.Getenv("TEST_DB_SSL_MODE"),
		},
		Migrator: Migrator{
			Env: os.Getenv("SERVER_ENV"),
		},
	}, nil
}

func (e *Environment) DatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		e.Database.Host,
		e.Database.Port,
		e.Database.User,
		e.Database.Password,
		e.Database.Name,
		e.Database.Ssl,
	)
}

func (e *Environment) TestDatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		e.TestDatabase.Host,
		e.TestDatabase.Port,
		e.TestDatabase.User,
		e.TestDatabase.Password,
		e.TestDatabase.Name,
		e.TestDatabase.Ssl,
	)
}

func getEnvPath() (string, error) {
	directory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filepath := searchup(directory, ".env")
	return filepath, nil
}

func searchup(dir string, filename string) string {
	if dir == "/" || dir == "" || dir == "." {
		return ""
	}

	if _, err := os.Stat(path.Join(dir, filename)); err == nil {
		return path.Join(dir, filename)
	}

	return searchup(path.Dir(dir), filename)
}
