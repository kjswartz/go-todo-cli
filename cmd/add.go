/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add todo item(s)",
	Long: `Add todo item(s) to the list. Accepts a list of todo items as a string slice to track. Can also specify the priority for the todo items.
	For example: 
		todo add -i "item 1, item 2, item 3".
		todo add -i "item 1" -p1.`,
	Run: addFunc,
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 3, "Priority of the todo item (1, 2, or 3)")
	addCmd.Flags().StringSliceP("items", "i", []string{}, "todo items to track, --iteams <setup app, setup db, etc>")
}

func addFunc(cmd *cobra.Command, args []string) {
	items, err := cmd.Flags().GetStringSlice("items")
	if err != nil {
		fmt.Println("Error getting items:", err)
		return
	}
	// if items not provided then ask user for input to provide a todo item
	if len(items) == 0 {
		fmt.Print("No todo items provided.\nPlease enter a todo item: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = strings.TrimSpace(input)
		if input != "" {
			items = append(items, input)
		} else {
			fmt.Println("No todo item entered")
			return
		}
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	for _, item := range items {
		insertSQL := `INSERT INTO todos (description, priority) VALUES (?, ?)`
		_, err := db.Exec(insertSQL, item, priority) // Default priority is 3
		if err != nil {
			fmt.Println("Error inserting item:", err)
			return
		}
		fmt.Println("Added todo:", item, "with priority", priority)
	}
}
