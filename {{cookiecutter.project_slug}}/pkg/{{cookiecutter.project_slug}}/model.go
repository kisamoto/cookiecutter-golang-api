package {{cookiecutter.project_slug}}

// Model is a base model template that can be configured with
// attributes as necessary
type Model struct {
	ID int `json:"id"`
}

// ModelService represents a set of methods for interacting
// with the Model. To be implemented by a store (e.g. PostgreSQL)
type ModelService interface {
	GetModels() ([]*Model, error)
	GetModelByID(modelID int) (*Model, error)
	UpdateModel(model *Model) (*Model, error)
	// Choose to "disable" over "delete"
	DisableModel(modelID int) error
	InsertModel(modle *Model) (*Model, error)
}
