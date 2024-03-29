package accesstoken

import (
	"bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

// return inteface
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetByID(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
