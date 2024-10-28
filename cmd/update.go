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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates todo item",
	Long: `Used to update a todo item. 
	For example: 
		todo update 1 -c to mark a todo item complete.
		todo update 1 -d "new description" to update the description of a todo item.
		todo update 1 -p2 to update the priority of a todo item.
		todo update 1 -r to remove/delete the todo item.`,
	Run: updateFunc,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateFunc(cmd *cobra.Command, args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	dbPath := filepath.Join(homeDir, "go", "data")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

}
