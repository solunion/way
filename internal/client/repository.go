package client

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
)

//goland:noinspection GoNameStartsWithPackageName
type ClientRepository struct {
	db *gorm.DB
}

func (t ClientRepository) FindAll(Clients *[]Client) (db *gorm.DB) {
	return t.db.Find(&Clients)
}

func (t ClientRepository) FindOne(Client *Client, id uuid.UUID) *gorm.DB {
	return t.db.Where("id = @id", sql.Named("id", id)).Take(&Client)
}

func (t ClientRepository) Create(Client *Client) (db *gorm.DB) {
	return t.db.Create(Client)
}

func (t ClientRepository) Update(Client *Client) (db *gorm.DB) {
	return t.db.Updates(&Client)
}

func (t ClientRepository) Save(Client *Client) (db *gorm.DB) {
	return t.Save(Client)
}

func (t ClientRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return t.db.Delete(id)
}

func NewClientRepository(DB *gorm.DB) internal.CRUDRepository[Client, uuid.UUID] {
	return &ClientRepository{db: DB}
}
