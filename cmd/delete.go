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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes todo item",
	Long: `Used to delete a todo item. 
	For example: 
		todo delete 1 to delete a todo item from the list.`,
	Run: deleteFunc,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteFunc(cmd *cobra.Command, args []string) {
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

	query := `DELETE FROM todos WHERE id = ?`
	_, err = db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting todo item:", err)
		return
	}

	fmt.Printf("Todo item %v has been deleted", id)
}
