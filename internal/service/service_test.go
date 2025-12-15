package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// stubRepo es nuestra implementación de prueba de UserRepo
type stubRepo struct {
	users          map[string]*User
	saveShouldFail bool
}

// NewStubRepo crea un nuevo stub con datos de prueba
func NewStubRepo() *stubRepo {
	return &stubRepo{
		users: map[string]*User{
			"juan@email.com":  {ID: 1, Email: "juan@email.com", Name: "Juan"},
			"maria@email.com": {ID: 2, Email: "maria@email.com", Name: "María"},
			"admin@email.com": {ID: 99, Email: "admin@email.com", Name: "Admin"},
		},
	}
}

// ByEmail implementa UserRepo.ByEmail
func (r *stubRepo) ByEmail(email string) (*User, error) {
	user, exists := r.users[email]
	if !exists {
		return nil, ErrNotFound
	}
	return user, nil
}

// Save implementa UserRepo.Save
func (r *stubRepo) Save(user *User) error {
	if r.saveShouldFail {
		return fmt.Errorf("simulated save error")
	}
	r.users[user.Email] = user
	return nil
}

// TestService_FindIDByEmail - Pruebas principales
func TestService_FindIDByEmail(t *testing.T) {
	// Creamos el stub con datos de prueba
	stub := NewStubRepo()
	service := New(stub)

	t.Run("Пользователь найден", func(t *testing.T) {
		// Arrange
		email := "juan@email.com"

		// Act
		id, err := service.FindIDByEmail(email)

		// Assert
		require.NoError(t, err, "No debería haber error")
		assert.Equal(t, int64(1), id, "ID debería ser 1")
	})

	t.Run("Пользователь не найден", func(t *testing.T) {
		// Arrange
		email := "noexiste@email.com"

		// Act
		id, err := service.FindIDByEmail(email)

		// Assert
		require.Error(t, err, "Debería haber error")
		assert.ErrorIs(t, err, ErrNotFound, "Error debería ser ErrNotFound")
		assert.Equal(t, int64(0), id, "ID debería ser 0 cuando hay error")
	})

	t.Run("Caso case-sensitive", func(t *testing.T) {
		// Nota: nuestro stub es case-sensitive
		id, err := service.FindIDByEmail("JUAN@email.com") // Mayúsculas
		require.Error(t, err)
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Equal(t, int64(0), id)
	})
}

// TestService_RegisterUser - Pruebas adicionales para registro
func TestService_RegisterUser(t *testing.T) {
	stub := NewStubRepo()
	service := New(stub)

	t.Run("Registro exitoso", func(t *testing.T) {
		email := "nuevo@email.com"
		name := "Nuevo Usuario"

		user, err := service.RegisterUser(email, name)

		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, name, user.Name)
		assert.Greater(t, user.ID, int64(0))

		// Verificamos que se guardó en el stub
		savedUser, err := stub.ByEmail(email)
		require.NoError(t, err)
		assert.Equal(t, user, savedUser)
	})

	t.Run("Email ya registrado", func(t *testing.T) {
		email := "juan@email.com" // Ya existe
		name := "Juan Duplicado"

		user, err := service.RegisterUser(email, name)

		require.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "already exists")
	})

	t.Run("Error al guardar", func(t *testing.T) {
		stub.saveShouldFail = true
		defer func() { stub.saveShouldFail = false }()

		user, err := service.RegisterUser("otro@email.com", "Test")

		require.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "failed to save")
	})
}

// Test con testify/mock (ejemplo avanzado)
func TestService_WithMock(t *testing.T) {
	// En lugar de un stub simple, podríamos usar testify/mock
	// para verificar interacciones específicas
	t.Run("Verificar que ByEmail fue llamado", func(t *testing.T) {
		// Este sería el enfoque con mocks reales
		// Por simplicidad, usamos nuestro stub pero verificamos estado
		stub := NewStubRepo()
		initialCount := len(stub.users)

		service := New(stub)
		_, err := service.RegisterUser("test@mock.com", "Mock User")

		require.NoError(t, err)
		// Verificamos que se agregó un usuario
		assert.Equal(t, initialCount+1, len(stub.users))
	})
}
