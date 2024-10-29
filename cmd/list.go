/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todo items",
	Long:  `List todo items by their assigned priority level from most ugrent p1 to least urgent p3.`,
	Run:   listFunc,
}

var showCompleted bool
var showAll bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&showCompleted, "completed", "c", false, "Show only completed todo items")
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all todo items (both completed and non-completed)")
}

func listFunc(cmd *cobra.Command, args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	dbPath := filepath.Join(homeDir, "go", "data", "todo.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	var query string
	var rows *sql.Rows

	if showAll {
		query = `SELECT id, description, priority, completed FROM todos ORDER BY priority ASC`
		rows, err = db.Query(query)
	} else if showCompleted {
		query = `SELECT id, description, priority, completed FROM todos WHERE completed = 1 ORDER BY priority ASC`
		rows, err = db.Query(query)
	} else {
		query = `SELECT id, description, priority, completed FROM todos WHERE completed = 0 ORDER BY priority ASC`
		rows, err = db.Query(query)
	}

	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var description string
		var priority int
		var completed bool

		err = rows.Scan(&id, &description, &priority, &completed)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		printTodoItem(priority, id, description, completed)
	}
}

// printTodoItem prints a todo item with its priority, ID, description, and completion status.
func printTodoItem(priority int, id int, description string, completed bool) {
	if completed {
		fmt.Printf("P%d | (%d) %s [c]\n", priority, id, description)
	} else {
		fmt.Printf("P%d | (%d) %s\n", priority, id, description)
	}
}
