package initializers

import (
	"22BSA10076_Backend/models"
)

func SyncDatabase() {
	// ✅ Now this works because `DB` is a `gorm.DB` instance
	DB.AutoMigrate(&models.User{})
}
