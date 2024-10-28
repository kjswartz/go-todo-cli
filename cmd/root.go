/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Simple command line todo app",
	Long:  `A simple command line todo app that allows you to add, prioritize (p1, p2, p3), complete or remove, and list todo items by priority.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initDB)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initDB()
}

func initDB() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	dataDir := filepath.Join(homeDir, "go", "data")
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		fmt.Println("Error creating data directory:", err)
		return
	}

	dbPath := filepath.Join(dataDir, "todo.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"description" TEXT,
		"priority" INTEGER,
		"completed" BOOLEAN DEFAULT FALSE
  );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
}