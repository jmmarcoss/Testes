package calculadora

import (
	"errors"
)

type Calculadora struct {
	valorA   float64
	valorB   float64
	operacao string
}

func NovaCalculadora(valorA, valorB float64, operacao string) *Calculadora {
	return &Calculadora{
		valorA:   valorA,
		valorB:   valorB,
		operacao: operacao,
	}
}

func (c *Calculadora) Soma() float64 {
	return c.valorA + c.valorB
}

func (c *Calculadora) Subtracao() float64 {
	return c.valorA - c.valorB
}

func (c *Calculadora) Multiplicacao() float64 {
	return c.valorA * c.valorB
}

func (c *Calculadora) Divisao() (float64, error) {
	if c.valorB == 0 {
		return 0, errors.New("erro: divisão por zero não é permitido")
	}
	return c.valorA / c.valorB, nil
}

func (c *Calculadora) RealizaCalculo() (float64, error) {
	switch c.operacao {
	case "+":
		return c.Soma(), nil
	case "-":
		return c.Subtracao(), nil
	case "*":
		return c.Multiplicacao(), nil
	case "/":
		return c.Divisao()
	default:
		return 0, errors.New("operação inválida")
	}
}
