package database

import (
	"backend/models"
	"backend/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Link{})

	if err != nil {
		fmt.Println(err)
		panic("Database Migration Failed!")
	}

	fmt.Println("Database Migration Successful!")
}
