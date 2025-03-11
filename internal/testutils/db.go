package testutils

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pujidjayanto/goginboilerplate/pkg/db"
	"github.com/stretchr/testify/assert"
)

func loadTestDatabaseDsn() (string, error) {
	envPath, err := getEnvPath()
	if err != nil || strings.TrimSpace(envPath) == "" {
		return "", fmt.Errorf("no .env file found")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return "", err
	}

	var (
		host     = os.Getenv("TEST_DB_HOST")
		port     = os.Getenv("TEST_DB_PORT")
		user     = os.Getenv("TEST_DB_USER")
		password = os.Getenv("TEST_DB_PASSWORD")
		dbName   = os.Getenv("TEST_DB_NAME")
	)

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		host,
		port,
		user,
		password,
		dbName,
	), nil
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

func NewTestDb(t *testing.T) *db.DatabaseHandler {
	dsn, err := loadTestDatabaseDsn()
	assert.NoError(t, err)

	db, err := db.InitDatabaseHandler(dsn)
	assert.NoError(t, err)

	t.Cleanup(func() {
		defer db.Close()

		// Get the directory of the current file
		_, currentFile, _, _ := runtime.Caller(0)
		currentDir := filepath.Dir(currentFile)

		// Construct the absolute path to teardown.sql
		teardownPath := filepath.Join(currentDir, "testdata", "teardown.sql")

		teardownScript, err := os.ReadFile(teardownPath)
		assert.NoError(t, err)

		err = db.GetDB(context.Background()).Exec(string(teardownScript)).Error
		assert.NoError(t, err)
	})

	return &db
}
