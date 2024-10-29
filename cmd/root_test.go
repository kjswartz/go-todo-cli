package cmd

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInitDB(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Error getting home directory: %v", err)
	}

	// Verify: Check if the database file was created
	testDbPath := filepath.Join(homeDir, "go", "data", "todo_test.db")
	dbPath = testDbPath // Override the global variable

	// Call the function to test
	initDB()

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
