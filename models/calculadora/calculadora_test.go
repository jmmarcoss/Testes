package calculadora

import (
	"testing"
)

// test Soma
func TestSoma(t *testing.T) {
	calc := NovaCalculadora(10, 5, "+")
	resultado, err := calc.RealizaCalculo()

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if resultado != 15 {
		t.Errorf("Esperado 15, mas obteve %.2f", resultado)
	}
}

// test Subtração
func TestSubtracao(t *testing.T) {
	calc := NovaCalculadora(10, 5, "-")
	resultado, err := calc.RealizaCalculo()

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if resultado != 5 {
		t.Errorf("Esperado 5, mas obteve %.2f", resultado)
	}
}

// test Multiplicação
func TestMultiplicacao(t *testing.T) {
	calc := NovaCalculadora(10, 5, "*")
	resultado, err := calc.RealizaCalculo()

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if resultado != 50 {
		t.Errorf("Esperado 50, mas obteve %.2f", resultado)
	}
}

// test Divisão
func TestDivisao(t *testing.T) {
	calc := NovaCalculadora(10, 5, "/")
	resultado, err := calc.RealizaCalculo()

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if resultado != 2 {
		t.Errorf("Esperado 2, mas obteve %.2f", resultado)
	}
}

// Teste para divisão por zero
func TestDivisaoPorZero(t *testing.T) {
	calc := NovaCalculadora(10, 0, "/")
	_, err := calc.RealizaCalculo()

	if err == nil {
		t.Errorf("Esperava um erro para divisão por zero")
	}
}

// Teste para operação inválida
func TestOperacaoInvalida(t *testing.T) {
	calc := NovaCalculadora(10, 5, "%")
	_, err := calc.RealizaCalculo()

	if err == nil {
		t.Errorf("Esperava um erro para operação inválida")
	}
}
