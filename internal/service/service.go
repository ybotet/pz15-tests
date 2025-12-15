package service

import "fmt"

// Service contiene la lógica de negocio
type Service struct {
	repo UserRepo
}

// New crea una nueva instancia de Service
func New(repo UserRepo) *Service {
	return &Service{repo: repo}
}

// FindIDByEmail busca un usuario por email y devuelve su ID
func (s *Service) FindIDByEmail(email string) (int64, error) {
	user, err := s.repo.ByEmail(email)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// RegisterUser registra un nuevo usuario (ejemplo adicional)
func (s *Service) RegisterUser(email, name string) (*User, error) {
	// Primero verificamos si el usuario ya existe
	existing, err := s.repo.ByEmail(email)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	// Si no existe, creamos uno nuevo
	// En un caso real, generaríamos un ID único
	newUser := &User{
		ID:    generateID(), // Función helper
		Email: email,
		Name:  name,
	}

	err = s.repo.Save(newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return newUser, nil
}

// Helper function para generar ID (simulación)
func generateID() int64 {
	// En producción usarías un generador real
	return int64(len("temporal") * 1000) // Solo para ejemplo
}
