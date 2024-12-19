package triangle

import "fmt"

type Triangulo struct {
	lado1 float64
	lado2 float64
	lado3 float64
}

func NovoTriangulo(lado1, lado2, lado3 float64) (*Triangulo, error) {
	if lado1 <= 0 || lado2 <= 0 || lado3 <= 0 {
		return nil, fmt.Errorf("todos os lados devem ser positivos")
	}

	return &Triangulo{
		lado1: lado1,
		lado2: lado2,
		lado3: lado3,
	}, nil
}

func (triangulo *Triangulo) Classificar() (string, error) {
	if !triangulo.EhTrianguloValido() {
		return "", fmt.Errorf("não é um triângulo válido")
	}

	if triangulo.lado1 == triangulo.lado2 && triangulo.lado2 == triangulo.lado3 {
		return "equilátero", nil
	}

	if triangulo.lado1 == triangulo.lado2 || triangulo.lado2 == triangulo.lado3 || triangulo.lado1 == triangulo.lado3 {
		return "isósceles", nil
	}

	return "escaleno", nil
}

func (triangulo *Triangulo) EhTrianguloValido() bool {
	return (triangulo.lado1+triangulo.lado2 > triangulo.lado3) &&
		(triangulo.lado1+triangulo.lado3 > triangulo.lado2) &&
		(triangulo.lado2+triangulo.lado3 > triangulo.lado1)
}

func (triangulo *Triangulo) calculaPerimetro() float64 {
	return triangulo.lado1 + triangulo.lado2 + triangulo.lado3
}
