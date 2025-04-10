/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates todo item",
	Long: `Used to update a todo item. 
	For example: 
		todo update 1 -c to mark a todo item complete.
		todo update 1 -d "new description" to update the description of a todo item.
		todo update 1 -p2 to update the priority of a todo item.`,
	Run: updateFunc,
}

var (
	markComplete bool
	description  string
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVarP(&markComplete, "complete", "c", false, "Mark the todo item as complete")
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "Update the description of the todo item")
}

func updateFunc(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID:", args[0])
		return
	}

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

	if markComplete {
		markCompleteFunc(db, id)
		return
	}

	if description != "" {
		updateDescriptionFunc(db, id, description)
		return
	}

	fmt.Println("No valid flags provided. Use -c to mark complete or -d to update description.")
}

func markCompleteFunc(db *sql.DB, id int) {
	query := `UPDATE todos SET completed = 1 WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Error updating todo item:", err)
		return
	}
	fmt.Println("Todo item", id, "marked as complete")
}

func updateDescriptionFunc(db *sql.DB, id int, description string) {
	query := `UPDATE todos SET description = ? WHERE id = ?`
	_, err := db.Exec(query, description, id)
	if err != nil {
		fmt.Println("Error updating todo description:", err)
		return
	}
	fmt.Println("Todo item", id, "description updated to:", description)
}
