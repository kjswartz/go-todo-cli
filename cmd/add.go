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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add todo item(s)",
	Long: `Add todo item(s) to the list. Will accept a single todo item as a string or a list of todo items as a string slice.
	For example: 
		todo add "item 1" "item 2" "item 3".
		todo add "item 1" -p1.`,
	Run: addFunc,
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 3, "Priority of the todo item (1, 2, or 3)")
}

var dbPath = filepath.Join(os.Getenv("HOME"), "go", "data", "todo.db")

func addFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("No todo items provided")
		return
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	for _, item := range args {
		insertSQL := `INSERT INTO todos (description, priority) VALUES (?, ?)`
		_, err := db.Exec(insertSQL, item, priority) // Default priority is 3
		if err != nil {
			fmt.Println("Error inserting item:", err)
			return
		}
		fmt.Println("Added todo:", item, "with priority", priority)
	}
}
