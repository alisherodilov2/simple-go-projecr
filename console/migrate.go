package console

import (
	"fmt"
	"log"

	"github.com/alisherodilov2/go-first/database"
	"github.com/alisherodilov2/go-first/models"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run all database migrations",
	Run: func(cmd *cobra.Command, args []string) {

		if err := database.DB.AutoMigrate(
			&models.Book{},
			&models.Products{},
			&models.Comments{},
			&models.User{},
		); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Migration complete ✅")
	},
}
