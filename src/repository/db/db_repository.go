// For using db interacting
package db

import (
	"bookstore_oauth-api/src/clients/cassandra"
	accesstoken "bookstore_oauth-api/src/domain/access_token"
	"errors"

	"github.com/Skele878/bookstore_utils-go/rest_errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetByID(string) (*accesstoken.AccessToken, rest_errors.RestErr)
	Create(accesstoken.AccessToken) rest_errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) rest_errors.RestErr
}
type dbRepository struct {
}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, rest_errors.RestErr) {
	var ressult accesstoken.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&ressult.AccessToken,
		&ressult.UserId,
		&ressult.ClientId,
		&ressult.Expires,
	); err != nil {

		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		return nil, rest_errors.NewInternalServerError("error when trying to get current id", errors.New("database error"))
	}

	return &ressult, nil
}

func (r *dbRepository) Create(at accesstoken.AccessToken) rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error when trying to save access token in database", err)
	}
	return nil
}
func (r *dbRepository) UpdateExpirationTime(at accesstoken.AccessToken) rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error when trying to update current resource", errors.New("database error"))
	}
	return nil
}
