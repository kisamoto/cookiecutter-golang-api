package {{cookiecutter.project_slug}}

import (
	"time"
)

// AccountCredentialsModel is a simple
type AccountCredentialsModel struct {
	ID                 int        `json:"id"`
	Email              string     `json:"email"`
	Password           []byte     `json:"password"`
	Enabled            bool       `json:"enabled"`
	CreateDateTime     *time.Time `json:"created_datetime"`
	LastActiveDateTime *time.Time `json:"last_active_datetime"`
}

type AccountProfileModel struct {
}
