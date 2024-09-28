package application

import "github.com/solunion/way/internal"

type Application struct {
	internal.UserModel `gorm:"embedded"`
	version            string
}
