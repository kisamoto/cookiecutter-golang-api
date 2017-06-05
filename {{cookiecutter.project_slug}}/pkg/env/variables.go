package env

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	DB           *sqlx.DB
	ListenerPort int
	Logger       *zap.Logger
	Debug        bool
)
