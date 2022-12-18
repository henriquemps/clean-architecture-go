package boot

import (
	"log"
	"project-name-here/boot/database"
	"project-name-here/entity"
)

func init() {

	migrate()
}

func migrate() {

	err := database.GetConnection().AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatalf("error run migrate \n details: %s", err)
	}
}
