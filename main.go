package main

import (
	"os"

	"github.com/alisherodilov2/go-first/console"
	"github.com/alisherodilov2/go-first/database"
	"github.com/alisherodilov2/go-first/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	enverr := godotenv.Load()
	if enverr != nil {
		return
	}

	database.Connect()

	r := gin.Default()
	routes.SetupRoutes(r)

	rootCmd := &cobra.Command{Use: "go-first"}

	rootCmd.AddCommand(console.MigrateCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	r.Run(":8080")
}
