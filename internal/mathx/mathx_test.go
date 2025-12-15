package mathx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSum_Table - Pruebas tabulares con testing estándar
func TestSum_Table(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"положительные", 2, 3, 5},
		{"отрицательные и положительные", 10, -5, 5},
		{"два отрицательных", -3, -7, -10},
		{"нулевые", 0, 0, 0},
		{"с нулем", 5, 0, 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Sum(c.a, c.b)
			if got != c.want {
				t.Errorf("Sum(%d, %d) = %d; want %d", c.a, c.b, got, c.want)
			}
		})
	}
}

// TestSum_WithTestify - Mismo test usando testify
func TestSum_WithTestify(t *testing.T) {
	assert.Equal(t, 5, Sum(2, 3))
	assert.Equal(t, -10, Sum(-3, -7))
	assert.Equal(t, 0, Sum(0, 0))
}

// TestDivide_OkAndError - Pruebas de Divide con testing estándar
func TestDivide_OkAndError(t *testing.T) {
	t.Run("Успешное разделение", func(t *testing.T) {
		result, err := Divide(10, 2)
		if err != nil {
			t.Fatal("No se esperaba error:", err)
		}
		if result != 5 {
			t.Errorf("Divide(10, 2) = %d; want 5", result)
		}
	})

	t.Run("División por cero", func(t *testing.T) {
		result, err := Divide(10, 0)
		if err == nil {
			t.Fatal("Se esperaba error 'divide by zero'")
		}
		if result != 0 {
			t.Errorf("Cuando hay error, resultado debe ser 0, got %d", result)
		}
		if err.Error() != "divide by zero" {
			t.Errorf("Mensaje de error incorrecto: %v", err)
		}
	})
}

// TestDivide_WithTestify - Usando testify
func TestDivide_WithTestify(t *testing.T) {
	// Caso exitoso
	result, err := Divide(10, 2)
	require.NoError(t, err) // Si hay error, test se detiene
	assert.Equal(t, 5, result)

	// Caso de error
	result, err = Divide(10, 0)
	require.Error(t, err)
	assert.Equal(t, 0, result)
	assert.Contains(t, err.Error(), "divide by zero")
}

// BenchmarkSum - Prueba de rendimiento (bonus)
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(123, 456)
	}
}

// TestPanic - Ejemplo de prueba para panic (bonus)
func TestPanic(t *testing.T) {
	// Función que hace panic
	panicFunc := func() {
		panic("error crítico")
	}

	// Verificamos que efectivamente hace panic
	assert.Panics(t, panicFunc, "La función debería hacer panic")

	// Verificamos el mensaje del panic
	assert.PanicsWithValue(t, "error crítico", panicFunc)
}
