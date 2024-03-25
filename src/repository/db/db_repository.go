// For using db interacting
package db

import (
	accesstoken "bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}
type dbRepository struct {
}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
