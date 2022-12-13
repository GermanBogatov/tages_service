package dao

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostgreSQLClient interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	//BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	//Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn2.CommandTag, error)
	//Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}
