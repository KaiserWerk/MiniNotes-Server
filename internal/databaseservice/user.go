package databaseservice

import (
	"encoding/base64"
	"github.com/KaiserWerk/mininotes-server/internal/entity"
	"github.com/KaiserWerk/mininotes-server/internal/helper"
)

func (ds *DatabaseService) GenerateUser() (*entity.User, error) {
	u := entity.User{
		Secret: base64.StdEncoding.EncodeToString(helper.GenerateSecret(15)),
	}

	result := ds.db.Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func (ds *DatabaseService) GetUser(id interface{}) (*entity.User, error) {
	var u entity.User
	result := ds.db.Find(&u, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func (ds DatabaseService) UpdateUser(user entity.User) error {
	result := ds.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}