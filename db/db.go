package db

import "context"

const (
	EnvDbName = "DB_NAME"
	EnvDbURI  = "DB_URI"
)

type Dropper interface {
	Drop(context.Context) error
}
