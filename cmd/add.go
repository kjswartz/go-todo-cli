/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add todo item",
	Long: `Add todo item to the list and pass in the -d description flag to provide a description of the todo item and -p for the priority.
	For example: 
		todo add -d "read a book" -p 1.`,
	Run: addFunc,
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 3, "Priority of the todo item (1, 2, or 3)")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Set the description of the todo item")
}

func addFunc(cmd *cobra.Command, args []string) {
	if description == "" {
		fmt.Println("No description flag provided.")
		return
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	insertSQL := `INSERT INTO todos (description, priority) VALUES (?, ?)`
	_, err = db.Exec(insertSQL, description, priority) // Default priority is 3
	if err != nil {
		fmt.Println("Error inserting item:", err)
		return
	}
	fmt.Println("Added todo:", description, "with priority", priority)
}
