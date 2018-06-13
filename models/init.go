package models

import "github.com/free-way/riverwaveMaps/helpers"

func RunMigration(){
	helpers.DB.AutoMigrate(
		&Map{},
	)
}
