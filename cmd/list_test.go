package cmd

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func TestListFunc(t *testing.T) {
	// Setup temporary database
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Error getting home directory: %v", err)
	}

	testDbPath := filepath.Join(homeDir, "go", "data", "todo_test.db")
	dbPath = testDbPath // Override the global variable

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}
	defer os.Remove(dbPath)
	defer db.Close()

	// Create table and insert test data
	_, err = db.Exec(`CREATE TABLE todos (id INTEGER PRIMARY KEY, description TEXT, priority INTEGER, completed BOOLEAN)`)
	if err != nil {
		t.Fatalf("Error creating table: %v", err)
	}

	_, err = db.Exec(`INSERT INTO todos (description, priority, completed) VALUES 
		('Test todo 1', 1, 0),
		('Test todo 2', 2, 1),
		('Test todo 3', 3, 0)`)
	if err != nil {
		t.Fatalf("Error inserting test data: %v", err)
	}

	tests := []struct {
		name           string
		showCompleted  bool
		showAll        bool
		expectedOutput string
	}{
		{
			name:           "Show all todos",
			showAll:        true,
			expectedOutput: "P1 | (1) Test todo 1\nP2 | (2) Test todo 2 [c]\nP3 | (3) Test todo 3\n",
		},
		{
			name:           "Show only completed todos",
			showCompleted:  true,
			expectedOutput: "P2 | (2) Test todo 2 [c]\n",
		},
		{
			name:           "Show only non-completed todos",
			expectedOutput: "P1 | (1) Test todo 1\nP3 | (3) Test todo 3\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			showCompleted = tt.showCompleted
			showAll = tt.showAll

			// Capture output
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			listFunc(&cobra.Command{}, []string{})

			w.Close()
			os.Stdout = old

			var buf [1024]byte
			n, _ := r.Read(buf[:])
			output := string(buf[:n])

			if output != tt.expectedOutput {
				t.Errorf("expected %q, got %q", tt.expectedOutput, output)
			}
		})
	}
}
