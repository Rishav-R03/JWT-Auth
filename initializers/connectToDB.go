package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ✅ Use `gorm.DB` instead of `sql.DB`
var DB *gorm.DB

func ConnectToDB() {
	// PostgreSQL connection string
	DSN := os.Getenv("DSN")
	var err error

	// ✅ Open DB using GORM
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to the database successfully!")
}
