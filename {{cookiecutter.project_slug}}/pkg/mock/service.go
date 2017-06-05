package mock

// Service implements a {{cookiecutter.project_slug}}.ModelService
// backed by a mock data store for testing
type Service struct {
}

// NewService returns a mock Service
func NewService() *Service {
	return &Service{}
}
