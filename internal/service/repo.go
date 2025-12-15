package service

import "fmt"

// User representa un usuario en el sistema
type User struct {
	ID    int64
	Email string
	Name  string // Agregamos Name para más realismo
}

// ErrNotFound es un error estándar para "no encontrado"
var ErrNotFound = fmt.Errorf("user not found")

// UserRepo es una interfaz para el repositorio de usuarios
// Define las operaciones que necesitamos sin implementación concreta
type UserRepo interface {
	ByEmail(email string) (*User, error)
	Save(user *User) error // Agregamos para ejemplos adicionales
}
