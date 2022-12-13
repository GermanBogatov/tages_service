package dao

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type ImageDAO struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
}

func NewImageStorage(client PostgreSQLClient) *ImageDAO {
	return &ImageDAO{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

func (s *ImageDAO) All(ctx context.Context) error {
	return nil
}

func (s *ImageDAO) Create(ctx context.Context) error {
	fmt.Println("CREATE DAO SHIIT")
	return nil
}
