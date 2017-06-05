package postgresql

import (
	"github.com/jmoiron/sqlx"

	"{{cookiecutter.repo}}/pkg/cache"
)

// Service implements go-api.ModelService
// backed by PostgreSQL
type Service struct {
	dbCache    cache.Cache
	create     *sqlx.NamedStmt
	selectByID *sqlx.Stmt
}

// NewService takes a database and cache.Cache creating a new Service
// with prepared statements.
func NewService(db *sqlx.DB, dbCache cache.Cache) (*Service, error) {
	createStmt, err := db.PrepareNamed(insertModel)
	if err != nil {
		return nil, err
	}
	selectByIDStmt, err := db.Preparex(selectModelByID)
	if err != nil {
		return nil, err
	}
	return &Service{
		dbCache:    dbCache,
		create:     createStmt,
		selectByID: selectByIDStmt,
	}, nil
}
