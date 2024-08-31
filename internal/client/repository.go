package client

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
)

//goland:noinspection GoNameStartsWithPackageName
type ClientRepository interface {
	internal.CRUDRepository[Client, uuid.UUID]
}

type clientRepository struct {
	db *gorm.DB
}

func (t clientRepository) FindAll(Clients *[]Client) (db *gorm.DB) {
	return t.db.Find(&Clients)
}

func (t clientRepository) FindOne(Client *Client, id uuid.UUID) *gorm.DB {
	return t.db.Where("id = @id", sql.Named("id", id)).Take(&Client)
}

func (t clientRepository) Create(Client *Client) (db *gorm.DB) {
	return t.db.Create(Client)
}

func (t clientRepository) Update(Client *Client) (db *gorm.DB) {
	return t.db.Updates(&Client)
}

func (t clientRepository) Save(Client *Client) (db *gorm.DB) {
	return t.Save(Client)
}

func (t clientRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return t.db.Delete(id)
}

func newClientRepository(DB *gorm.DB) *clientRepository {
	return &clientRepository{db: DB}
}
