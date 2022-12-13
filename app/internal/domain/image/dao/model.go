package dao

import (
	"database/sql"
)

type ImageStorage struct {
	ID        string
	Name      string
	Size      uint64
	Bytes     []byte
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
}
