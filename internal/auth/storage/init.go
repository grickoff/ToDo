package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type AuthStorage struct {
	repo         *sqlx.DB
	queryBuilder sq.StatementBuilderType
}

func NewAuthStorage(repo *sqlx.DB) *AuthStorage {
	return &AuthStorage{
		repo:         repo,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
