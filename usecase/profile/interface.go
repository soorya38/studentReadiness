package profile

type Repository interface {
	Writer
}

type Writer interface {
	Create() error
}

type Usecase interface {
	CreateProfile() error
}