package cmd

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInitDB(t *testing.T) {
	// Setup: Create a temporary directory for the test
	tempDir := t.TempDir()
	originalHomeDir := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHomeDir)
	os.Setenv("HOME", tempDir)

	// Call the function to test
	initDB()

	// Verify: Check if the database file was created
	dbPath := filepath.Join(tempDir, "go", "data", "todo_test.db")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Fatalf("Expected database file to be created at %s, but it does not exist", dbPath)
	}

	// Verify: Check if the table was created
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='todos';")
	var tableName string
	if err := row.Scan(&tableName); err != nil {
		t.Fatalf("Error querying table: %v", err)
	}
	if tableName != "todos" {
		t.Fatalf("Expected table 'todos' to be created, but it was not")
	}
}
