package profile

type Service struct {
	repo Repository
}

func NewService(s Repository) *Service {
	return &Service{
		repo: s,
	}
}

func (s *Service) CreateProfile() error {
	if err := s.repo.Create(); err != nil {
		return err
	}
	return nil
}
