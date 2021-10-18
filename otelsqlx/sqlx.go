package otelsqlx

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

// Connect to a database and verify with a ping.
func Connect(driverName, dataSourceName string, opts ...otelsql.Option) (*sqlx.DB, error) {
	return ConnectContext(context.Background(), driverName, dataSourceName, opts...)
}

// ConnectContext to a database and verify with a ping.
func ConnectContext(
	ctx context.Context, driverName, dataSourceName string, opts ...otelsql.Option,
) (*sqlx.DB, error) {
	db, err := Open(driverName, dataSourceName, opts...)
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

// MustConnect connects to a database and panics on error.
func MustConnect(driverName, dataSourceName string, opts ...otelsql.Option) *sqlx.DB {
	db, err := Connect(driverName, dataSourceName, opts...)
	if err != nil {
		panic(err)
	}
	return db
}

// Open is the same as sql.Open, but returns an *sqlx.DB instead.
func Open(driverName, dataSourceName string, opts ...otelsql.Option) (*sqlx.DB, error) {
	db, err := otelsql.Open(driverName, dataSourceName, opts...)
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(db, driverName), nil
}

// MustOpen is the same as sql.Open, but returns an *sqlx.DB instead and panics on error.
func MustOpen(driverName, dataSourceName string, opts ...otelsql.Option) *sqlx.DB {
	db, err := Open(driverName, dataSourceName, opts...)
	if err != nil {
		panic(err)
	}
	return db
}
