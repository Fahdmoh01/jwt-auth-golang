package initializers

import (
	"jwt-auth-golang/models"
)

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}