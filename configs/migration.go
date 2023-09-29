package configs

import "unjuk_prakerja/models/users"

func initMigrate() {
	DB.AutoMigrate(&users.User{})
}
