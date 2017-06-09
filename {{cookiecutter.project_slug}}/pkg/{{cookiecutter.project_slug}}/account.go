package {{cookiecutter.project_slug}}

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptHashCost = 22
)

// AccountCredentialsModel is a simple representation
// credentials used to authenticate an account. 
type AccountCredentialsModel struct {
	ID                 			int        `json:"id"`
	Email              			string     `json:"email"`
	PasswordHash       			[]byte     `json:"password"`
	Enabled            			bool       `json:"enabled"`
	CreateDateTime     			*time.Time `json:"created_datetime"`
	LastActiveDateTime 			*time.Time `json:"last_active_datetime"`
	PasswordUpdatedDateTime *time.Time `json:"password_updated_datetime"`
}

func (acc *AccountCredentialsModel) PasswordMatches(password []byte) bool {
	return bcrypt.CompareHashAndPassword(acc.PasswordHash, password) == nil
}

func (acc *AccountCredentialsModel) HashAndSetPassword(password []byte) error {
	passwordHash, err := bcrypt.GenerateFromPassword(password, bcryptHashCost)
	if err != nil {
		return err
	}
	acc.PasswordHash = passwordHash
	return nil
}

type AccountProfileModel struct {
}

// AccountService interacts with both AccountCredentials and
// AccountProfile models. 
type AccountService interface {
	CreateAccountCredentials(acc *AccountCredentialsModel) (*AccountCredentialsModel, error) 
	GetAccountCredentialsByEmail(email string) (*AccountCredentialsModel, error) 
	UpdateAccountCredentials(acc *AccountCredentialsModel) (*AccountCredentialsModel, error) 
	Disable(accountID int) error
}
