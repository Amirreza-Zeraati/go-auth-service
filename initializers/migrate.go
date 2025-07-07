package initializers

import "jwt/models"

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
