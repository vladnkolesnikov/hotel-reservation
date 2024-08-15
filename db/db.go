package db

import "context"

const EnvDBName = "DB_NAME"

type Dropper interface {
	Drop(context.Context) error
}
