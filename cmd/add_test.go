package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func TestAddFunc(t *testing.T) {
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
	defer db.Close()

	// Ensure the todos table exists
	// createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
	//       id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	//       description TEXT,
	//       priority INTEGER,
	//       completed BOOLEAN DEFAULT FALSE
	//   );`
	// _, err = db.Exec(createTableSQL)
	// if err != nil {
	// 	t.Fatalf("Error creating table: %v", err)
	// }

	tests := []struct {
		name     string
		args     []string
		priority int
		wantErr  bool
	}{
		{"Add single item", []string{"Buy milk"}, 3, false},
		{"Add multiple items", []string{"Read book", "Go for a walk"}, 2, false},
		{"No items", []string{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear the table before each test
			_, err := db.Exec("DELETE FROM todos")
			if err != nil {
				t.Fatalf("Error clearing table: %v", err)
			}

			cmd := &cobra.Command{}
			cmd.Flags().IntP("priority", "p", tt.priority, "Priority of the todo item (1, 2, or 3)")
			cmd.Flags().Set("priority", fmt.Sprintf("%d", tt.priority))
			addFunc(cmd, tt.args)

			if tt.wantErr {
				// Check if error message was printed
				if len(tt.args) == 0 {
					t.Log("Expected error for no todo items provided")
				}
			} else {
				// Verify items were added to the database
				for _, item := range tt.args {
					var count int
					query := `SELECT COUNT(*) FROM todos WHERE description = ? AND priority = ?`
					err := db.QueryRow(query, item, tt.priority).Scan(&count)
					if err != nil {
						t.Fatalf("Error querying database: %v", err)
					}
					if count != 1 {
						t.Errorf("Expected 1 record for item %q with priority %d, got %d", item, tt.priority, count)
					}
				}
			}
		})
	}
}
