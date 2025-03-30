package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ✅ Use `gorm.DB` instead of `sql.DB`
var DB *gorm.DB

func ConnectToDB() {
	// PostgreSQL connection string
	dsn := "host=localhost user=postgres password=root123 dbname=File_Mng port=5432 sslmode=disable"
	var err error

	// ✅ Open DB using GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to the database successfully!")
}
