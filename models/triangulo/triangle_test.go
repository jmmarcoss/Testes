package triangle

import (
	"testing"
)

func TestTriangulo(t *testing.T) {

	// Triângulo escaleno válido
	t.Run("Triângulo Escaleno Válido", func(t *testing.T) {
		triangulo, err := NovoTriangulo(3, 4, 5)
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		tipo, err := triangulo.Classificar()
		if err != nil {
			t.Errorf("Erro na classificação: %v", err)
		}

		if tipo != "escaleno" {
			t.Errorf("Esperado escaleno, obtido %s", tipo)
		}
	})

	// Triângulo isósceles válido
	t.Run("Triângulo Isósceles Válido", func(t *testing.T) {
		triangulo, err := NovoTriangulo(5, 5, 7)
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		tipo, err := triangulo.Classificar()
		if err != nil {
			t.Errorf("Erro na classificação: %v", err)
		}

		if tipo != "isósceles" {
			t.Errorf("Esperado isósceles, obtido %s", tipo)
		}
	})

	// Triângulo equilátero válido
	t.Run("Triângulo Equilátero Válido", func(t *testing.T) {
		triangulo, err := NovoTriangulo(6, 6, 6)
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		tipo, err := triangulo.Classificar()
		if err != nil {
			t.Errorf("Erro na classificação: %v", err)
		}

		if tipo != "equilátero" {
			t.Errorf("Esperado equilátero, obtido %s", tipo)
		}
	})

	// Lado informado pelo usuário com valor zero
	t.Run("Lado com Valor Zero", func(t *testing.T) {
		_, err := NovoTriangulo(0, 4, 5)
		if err == nil {
			t.Errorf("Esperado erro para lado zero")
		}
	})

	// Lado informado pelo usuário com valor negativo
	t.Run("Lado com Valor Negativo", func(t *testing.T) {
		_, err := NovoTriangulo(-3, 4, 5)
		if err == nil {
			t.Errorf("Esperado erro para lado negativo")
		}
	})

	// Situação onde a soma de 2 lados é igual ao terceiro lado
	t.Run("Soma de Lados Igual ao Terceiro", func(t *testing.T) {
		triangulo, err := NovoTriangulo(3, 4, 7)
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		if triangulo.EhTrianguloValido() {
			t.Errorf("Esperado triângulo inválido")
		}
	})
}

func TestCalculaPerimetro(t *testing.T) {
	tests := []struct {
		nome      string
		triangulo Triangulo
		esperado  float64
	}{
		{
			nome:      "Triângulo equilátero",
			triangulo: Triangulo{lado1: 3, lado2: 3, lado3: 3},
			esperado:  9,
		},
		{
			nome:      "Triângulo isósceles",
			triangulo: Triangulo{lado1: 4, lado2: 4, lado3: 6},
			esperado:  14,
		},
		{
			nome:      "Triângulo escaleno",
			triangulo: Triangulo{lado1: 5, lado2: 7, lado3: 10},
			esperado:  22,
		},
		{
			nome:      "Triângulo com valores decimais",
			triangulo: Triangulo{lado1: 2.5, lado2: 3.5, lado3: 4},
			esperado:  10,
		},
	}

	for _, test := range tests {
		t.Run(test.nome, func(t *testing.T) {
			perimetro := test.triangulo.calculaPerimetro()
			if perimetro != test.esperado {
				t.Errorf("Esperado %.2f, mas obteve %.2f", test.esperado, perimetro)
			}
		})
	}
}
