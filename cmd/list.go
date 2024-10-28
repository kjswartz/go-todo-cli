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

func init() {
	rootCmd.AddCommand(listCmd)
}

func listFunc(cmd *cobra.Command, args []string) {
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

	// Query the database for all todo items ordered by priority
	query := `SELECT item, priority FROM todos ORDER BY priority ASC`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var item string
		var priority int
		err = rows.Scan(&item, &priority)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("P%d | %s\n", priority, item)
	}
}
