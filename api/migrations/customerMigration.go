package migrations

import (
	conn "interview/api/connections"
	"interview/api/models"
)


func MigrateCustomer() {
	conn.DB.AutoMigrate(&models.Customer{})
}