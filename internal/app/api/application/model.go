package application

import "github.com/solunion/way/internal"

type Application struct {
	internal.UserModel `bun:",extend"`
	version            string `bun:"type:text,notnull"`
}
