package repositories

type Repository[Model any] interface {
	FindById(id string) (*Model, error)
	Delete(Model) error
	Update(Model) error
	List() ([]Model, error)
}
